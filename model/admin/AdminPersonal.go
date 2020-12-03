package admin

import (
	"linea-esperanza-restapi/model/personal"
)

type AdminPersonal struct {
	personal.Personal
	rolePersonal string `json:"rolePersonal,omitempty"`
}

type AdminPersonalService interface {
	GetAll() ([]AdminPersonal, error)
	Save(adminPersonal *AdminPersonal) error
}

var adminService AdminPersonalService

func (adminPersonal *AdminPersonal) Save() error {
	return adminService.Save(adminPersonal)
}

func (adminPersonal *AdminPersonal) GetAll() ([]AdminPersonal, error) {
	return adminService.GetAll()
}
