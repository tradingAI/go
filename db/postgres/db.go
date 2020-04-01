package postgres

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

// CreateTables create all tables
func CreateTables(db *gorm.DB, values ...interface{}) (err error) {
	errs := db.CreateTable(values...).GetErrors()
	if len(errs) != 0 {
		for _, err := range errs {
			glog.Error(err)
		}
		return
	}

	return
}

// DropTables drop all tables
func DropTables(db *gorm.DB, values ...interface{}) (err error) {
	errs := db.DropTableIfExists(values...).GetErrors()
	if len(errs) != 0 {
		for _, err := range errs {
			glog.Error(err)
		}
		return
	}

	return
}

// ResetTables drop and create tables
func ResetTables(db *gorm.DB, values ...interface{}) (err error) {
	if err = DropTables(db, values...); err != nil {
		glog.Error(err)
		return
	}

	if err = CreateTables(db, values...); err != nil {
		glog.Error(err)
		return
	}

	return
}
