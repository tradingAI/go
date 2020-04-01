package redis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/glog"
	err2 "github.com/tradingAI/go/error"
)

type RedisConf struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

func (c *RedisConf) Validate() (err error) {
	if c.Host == "" {
		err = err2.ErrEmptyRedisHost
		glog.Error(err)
		return
	}

	if c.Port <= 1024 || c.Port >= 65535 {
		err = err2.ErrInvaldiRedisPort
		glog.Error(err)
		return
	}

	return
}

func NewRedisClient(conf RedisConf) (c *redis.Conn, err error) {
	client, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", conf.Host, conf.Port))
	if err != nil {
		glog.Error(err)
		return
	}

	c = &client

	return
}
