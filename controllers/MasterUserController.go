package controllers

import (
	"govue/helpers"

	knot "github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
)

type MasterUserController struct {
	Config *helpers.Configuration
}

func (c *MasterUserController) Default(r *knot.WebContext) interface{} {
	r.Config.ViewName = "admin/master_user.html"
	data := toolkit.M{}
	data["menu"] = "master_user"
	data["message"] = "this is message from go"
	return data
}
