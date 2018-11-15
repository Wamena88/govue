package models

import (
	"github.com/eaciit/orm"
	"gopkg.in/mgo.v2/bson"
)

type MasterUserModel struct {
	orm.ModelBase `bson:"-" json:"-"`
	Id            bson.ObjectId `bson:"_id"  json:"_id" `
	Username      string        `json:"username"`
	Password      string        `json:"password"`
	Name          string        `json:"name"`
	Email         string        `json:"email"`
}
