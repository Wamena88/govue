package main

import (
	"govue/controllers"
	"govue/helpers"
	"net/http"
	"os"

	"github.com/eaciit/knot/knot.v1"
	tk "github.com/eaciit/toolkit"
)

var (
	appPath = func() string {
		d, _ := os.Getwd()
		return d
	}()
)

func main() {
	app := knot.NewApp("WebVue")
	confInit := helpers.Configuration{}
	confInit.Connection = helpers.CreateConn()
	conf := helpers.New(&confInit)

	app.ViewsPath = appPath + tk.PathSeparator + "views" + tk.PathSeparator
	app.Static("static", appPath+tk.PathSeparator+"assets"+tk.PathSeparator)
	app.Static("files", appPath+tk.PathSeparator+"uploads"+tk.PathSeparator)
	app.Static("data", appPath+tk.PathSeparator+"data"+tk.PathSeparator)
	app.Register(&controllers.DashboardController{Config: conf})
	app.Register(&controllers.MasterUserController{Config: conf})
	app.LayoutTemplate = "_template.html"
	app.DefaultOutputType = knot.OutputTemplate
	knot.RegisterApp(app)

	otherRoutes := map[string]knot.FnContent{
		"/": func(r *knot.WebContext) interface{} {
			http.Redirect(r.Writer, r.Request, "/dashboard/index", http.StatusTemporaryRedirect)
			return true
		},
	}

	knot.StartAppWithFn(app, helpers.Host+helpers.PortServer, otherRoutes)

}
