package postgres

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	err2 "github.com/tradingAI/go/error"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBConf struct {
	Database     string        `yaml:"database"`
	Username     string        `yaml:"username"`
	Password     string        `yaml:"password"`
	Port         int           `yaml:"port"`
	Host         string        `yaml:"host"`
	Reset        bool          `yaml:"reset"`
	ReconnectSec time.Duration `yaml:"reconnect_sec"`
}

func (c *DBConf) Validate() (err error) {
	if c.Host == "" {
		err = err2.ErrEmptyDBHost
		glog.Error(err)
		return
	}

	if c.Port <= 1024 || c.Port >= 65535 {
		err = err2.ErrInvalidDBPort
		glog.Error(err)
		return
	}

	if c.Username == "" {
		err = err2.ErrEmptyDBUsername
		glog.Error(err)
		return
	}

	if c.Password == "" {
		err = err2.ErrEmptyDBPassword
		glog.Error(err)
		return
	}

	if c.Database == "" {
		err = err2.ErrEmptyDBDatabase
		glog.Error(err)
		return
	}

	return
}

func NewPostgreSQL(conf DBConf) (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
			conf.Host,
			conf.Port,
			conf.Username,
			conf.Database,
			conf.Password))

	if err != nil {
		glog.Error(err)
		return
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return
}
