package controllers

import (
	"webvue/helpers"

	knot "github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
)

type DashboardController struct {
	Config *helpers.Configuration
}

func (c *DashboardController) Index(r *knot.WebContext) interface{} {
	r.Config.ViewName = "admin/dashboard.html"
	data := toolkit.M{}
	data["menu"] = "dashboard"
	data["message"] = "this is message from go"
	return data
}
