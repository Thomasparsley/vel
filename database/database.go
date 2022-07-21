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
	return instance
}
