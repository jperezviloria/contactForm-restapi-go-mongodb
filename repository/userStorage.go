package repository

import (
	"gopkg.in/mgo.v2"
	"linea-esperanza-restapi/model/admin"
)

type UserStorage struct {
	collection *mgo.Collection
	context    *Context
}

func (u UserStorage) GetAll() ([]admin.AdminPersonal, error) {
	panic("implement me")
}

func (u UserStorage) Save(adminPersonal *admin.AdminPersonal) error {
	panic("implement me")
}
