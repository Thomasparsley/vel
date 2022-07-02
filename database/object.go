package database

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Thomasparsley/vel/structs/optional"
	"github.com/Thomasparsley/vel/structs/result"
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
func (o Object[ID, O]) First() result.Result[optional.Optional[O]] {
	var data O

	response := o.query.First(&data)
	if response.Error != nil {
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			return result.Ok(optional.None[O]())
		}

		return result.Error[optional.Optional[O]](response.Error)
	}

	return result.Ok(optional.Some(data))
}

// Find find records that match given conditions
func (o Object[ID, O]) Find() result.Result[[]O] {
	var data []O

	response := o.query.Find(&data)
	if response.Error != nil {
		return result.Error[[]O](response.Error)
	}

	return result.Ok(data)
}

// Exists is predecate for object existence
func (o Object[ID, O]) Exists() result.Result[bool] {
	var data O
	var zeroValueID ID

	response := o.query.First(&data)
	if response.Error != nil {
		return result.Error[bool](response.Error)
	}

	return result.Ok(data.PK() != zeroValueID)
}

func (o Object[ID, O]) Count() result.Result[int64] {
	var data int64

	response := o.query.Count(&data)
	if response.Error != nil {
		return result.Error[int64](response.Error)
	}

	return result.Ok(data)
}

// Create insert the value into database
func (o Object[ID, O]) Create(data O) result.Result[O] {
	response := o.query.Create(&data)
	if response.Error != nil {
		return result.Error[O](response.Error)
	}

	return result.Ok(data)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (o Object[ID, O]) FirstOrCreate(data O) result.Result[optional.Optional[O]] {
	response := o.query.FirstOrCreate(&data)
	if response.Error != nil {
		return result.Error[optional.Optional[O]](response.Error)
	}

	return result.Ok(optional.Some(data))
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (o Object[ID, O]) Save(data O) result.Result[O] {
	response := o.query.Save(&data)
	if response.Error != nil {
		return result.Error[O](response.Error)
	}

	return result.Ok(data)
}
