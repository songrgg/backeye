package model

import (
	"github.com/jinzhu/gorm"
	"github.com/songrgg/backeye/std"
)

var (
	db *gorm.DB
)

// InitModel initialize the MySQL
func InitModel(config std.ConfigMySQL) {
	db = std.CreateDB(config)
	std.LogInfoLn("start init mysql model")

	db.AutoMigrate(&AssertionResult{}, &Watcher{}, &WatcherSnapshot{})
	std.LogInfoLn("end init mysql model")
}

func DB() *gorm.DB {
	return db
}
