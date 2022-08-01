package database

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	mutex    sync.Mutex
	instance *gorm.DB = nil
)

type Config struct {
	Host         string
	User         string
	Password     string
	DatabaseName string
	Port         uint
	SSL          bool
	TimeZone     string
}

func (c Config) dsn() string {
	sslMode := "disabled"
	if c.SSL {
		sslMode = "enabled"
	}

	return fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		c.User,
		c.Password,
		c.DatabaseName,
		c.Port,
		sslMode,
		c.TimeZone,
	)
}

func Initialize(config Config) {
	var err error

	mutex.Lock()
	defer mutex.Unlock()

	instance, err = gorm.Open(
		postgres.Open(config.dsn()),
		&gorm.Config{SkipDefaultTransaction: true},
	)
	if err != nil {
		panic(err)
	}
}

func Get() *gorm.DB {
	if instance == nil {
		panic("database connection instace wasnt initialize")
	}

	return instance
}
