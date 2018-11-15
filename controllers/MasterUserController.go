package controllers

import (
	daos "govue/dao"
	"govue/helpers"
	"govue/models"

	knot "github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
)

type MasterUserController struct {
	Config *helpers.Configuration
}

func (c *MasterUserController) populateData() ([]models.MasterUserModel, bool) {
	var ret = []models.MasterUserModel{}
	var dao = new(daos.MasterUserDao)
	dao.Conn = c.Config.Connection
	ret = dao.FindAll()
	if ret == nil {
		return ret, false
	}
	return ret, true
}
func (c *MasterUserController) Default(r *knot.WebContext) interface{} {
	r.Config.ViewName = "admin/master_user.html"
	data := toolkit.M{}
	data["menu"] = "master_user"
	m, isValid := c.populateData()
	if isValid {
		data["populate"] = m
	} else {
		data["populate"] = nil
	}
	return data
}
