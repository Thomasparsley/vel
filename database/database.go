package database

import (
	"sync"

	"gorm.io/gorm"
)

var (
	mutex    sync.Mutex
	instance *gorm.DB = nil
)

func Initialize() {
	mutex.Lock()
	defer mutex.Unlock()
}

func Get() *gorm.DB {
	if instance == nil {
		panic("database connection instace wasnt initialize")
	}

	return instance
}
