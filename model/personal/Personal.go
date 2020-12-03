package personal

import "go.mongodb.org/mongo-driver/bson/primitive"

type Personal struct {
	IdPersonal       primitive.ObjectID `json:"_id,omitempty"`
	NamePersonal     string             `json:"namePersonal,omitempty"`
	CountryPersonal  string             `json:"countryPersonal,omitempty"`
	PasswordPersonal string             `json:"passwordPersonal,omitempty"`
}
