package dao

import (
	"github.com/gocraft/dbr"
	"github.com/liangzuobin/gomocktest/double/model"
)

// ModelByID find a model by its id
func ModelByID(sr dbr.SessionRunner, id int64) (*model.Model, error) {
	var m model.Model
	err := sr.Select("*").From("model").Where("id = ?", id).LoadStruct(&m)
	return &m, err
}
