package hashids

import (
	"sync"

	"github.com/speps/go-hashids/v2"
)

var (
	mutex    sync.Mutex
	config   *hashids.HashIDData = nil
	instance *hashids.HashID     = nil
)

func Initialize(salt string, minLength uint) error {
	var err error

	mutex.Lock()
	defer mutex.Unlock()

	if config == nil && instance == nil {
		config = hashids.NewData()
		config.Salt = salt
		config.MinLength = int(minLength)

		instance, err = hashids.NewWithData(config)
	}

	return err
}

func Get() *hashids.HashID {
	if instance == nil {
		panic("hashids instace wasnt initialize")
	}

	return instance
}
