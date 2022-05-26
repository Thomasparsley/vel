package vel

import "gorm.io/gorm"

type ObjectDefinition interface {
	PK() int
	TableName() string
}

type Object[O ObjectDefinition] struct {
	query *gorm.DB
}

func NewObject[O ObjectDefinition](db *gorm.DB) *Object[O] {
	var o O

	return &Object[O]{
		query: db.Table(o.TableName()),
	}
}

// Where add conditions
func (o *Object[O]) Where(query interface{}, args ...interface{}) *Object[O] {
	o.query = o.query.Where(query, args...)
	return o
}

// Joins specify Joins conditions
func (o *Object[O]) Joins(query string, args ...interface{}) *Object[O] {
	o.query = o.query.Joins(query, args...)
	return o
}

// Preload preload associations with given conditions
func (o *Object[O]) Preload(query string, args ...interface{}) *Object[O] {
	o.query = o.query.Preload(query, args...)
	return o
}

// First find first record that match given conditions, order by primary key
func (o *Object[O]) First() *O {
	var data O

	o.query.First(&data)
	if data.PK() == 0 {
		return nil
	}

	return &data
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (o *Object[O]) FirstOrCreate() *O {
	var data O

	o.query.FirstOrCreate(&data)

	return &data
}

// Find find records that match given conditions
func (o *Object[O]) Find() []O {
	var data []O

	o.query.Find(&data)

	return data
}

func (o *Object[O]) Exists() bool {
	var data O

	o.query.First(&data)

	return data.PK() != 0
}

func (o *Object[O]) Count() int64 {
	var data int64

	o.query.Count(&data)

	return data
}
