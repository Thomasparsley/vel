package database

import (
	"github.com/Thomasparsley/vel/optional"
	"gorm.io/gorm"
)

type ComperableID interface {
	string | uint64
}

type ObjectDefinition[ID ComperableID] interface {
	PK() ID
	TableName() string
}

type Object[ID ComperableID, O ObjectDefinition[ID]] struct {
	query *gorm.DB
}

func NewObject[ID ComperableID, O ObjectDefinition[ID]](db *gorm.DB) Object[ID, O] {
	var o O

	return Object[ID, O]{
		query: db.Table(o.TableName()),
	}
}

// Where add conditions
func (o Object[ID, O]) Where(query interface{}, args ...interface{}) Object[ID, O] {
	o.query = o.query.Where(query, args...)
	return o
}

// Joins specify Joins conditions
func (o Object[ID, O]) Joins(query string, args ...interface{}) Object[ID, O] {
	o.query = o.query.Joins(query, args...)
	return o
}

// Preload preload associations with given conditions
func (o Object[ID, O]) Preload(query string, args ...interface{}) Object[ID, O] {
	o.query = o.query.Preload(query, args...)
	return o
}

// First find first record that match given conditions, order by primary key
func (o Object[ID, O]) First() optional.Optional[O] {
	var data O
	var zeroValueID ID

	o.query.First(&data)
	if data.PK() == zeroValueID {
		return optional.None[O]()
	}

	return optional.Some(data)
}

// Find find records that match given conditions
func (o Object[ID, O]) Find() []O {
	var data []O

	o.query.Find(&data)

	return data
}

// Exists is predecate for object existence
func (o Object[ID, O]) Exists() bool {
	var data O
	var zeroValueID ID

	o.query.First(&data)

	return data.PK() != zeroValueID
}

func (o Object[ID, O]) Count() int64 {
	var data int64

	o.query.Count(&data)

	return data
}

// Create insert the value into database
func (o Object[ID, O]) Create(data O) O {
	o.query.Create(&data)
	return data
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (o Object[ID, O]) FirstOrCreate(data O) optional.Optional[O] {
	o.query.FirstOrCreate(&data)

	return optional.Some(data)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (o Object[ID, O]) Save(data O) O {
	o.query.Save(&data)
	return data
}
