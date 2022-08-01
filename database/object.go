package database

import (
	"errors"

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
func (o Object[ID, O]) First() (*O, error) {
	var data O

	response := o.query.First(&data)
	if response.Error != nil {
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, response.Error
	}

	return &data, nil
}

// Find find records that match given conditions
func (o Object[ID, O]) Find() ([]O, error) {
	var data []O

	response := o.query.Find(&data)
	if response.Error != nil {
		return data, response.Error
	}

	return data, nil
}

// Exists is predecate for object existence
func (o Object[ID, O]) Exists() (bool, error) {
	var data O
	var zeroValueID ID

	response := o.query.First(&data)
	if errors.Is(response.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if response.Error != nil {
		return false, response.Error
	}

	return data.PK() != zeroValueID, nil
}

func (o Object[ID, O]) Count() (int64, error) {
	var data int64

	response := o.query.Count(&data)
	if response.Error != nil {
		return 0, response.Error
	}

	return data, nil
}

// Create insert the value into database
func (o Object[ID, O]) Create(data O) (O, error) {
	response := o.query.Create(&data)
	if response.Error != nil {
		return data, response.Error
	}

	return data, nil
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (o Object[ID, O]) FirstOrCreate(data O) (*O, error) {
	response := o.query.FirstOrCreate(&data)
	if response.Error != nil {
		return nil, response.Error
	}

	return &data, nil
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (o Object[ID, O]) Save(data O) (O, error) {
	response := o.query.Save(&data)
	if response.Error != nil {
		return data, response.Error
	}

	return data, nil
}
