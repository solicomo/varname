package main

import (
	"net/http"
	"varname/app/controllers"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "app/views/simple",
		Extensions: []string{".tmpl", ".html"},
	}))

	m.Get("/search", func(req *http.Request, r render.Render) {
		qs := req.FormValue("q")
		qs = controllers.Normalize(qs)

		data := make(map[string]interface{})
		data["Title"] = qs
		data["PageClass"] = "search"
		data["SearchString"] = qs
		data["SearchRelates"] = []string{qs, qs, qs, qs}
		data["SearchList"], err = constrollers.Search(0, 10, qs)

		r.HTML(200, "search", data)
	})

	m.Get("/e/:entry", func(r render.Render, params martini.Params) {
		r.HTML(200, "entry", map[string]interface{}{"entry": params["entry"]})
	})

	m.Get("/(?P<page>(about|terms))", func(r render.Render, params martini.Params) {
		r.HTML(200, "page", map[string]interface{}{
			"PageHead":    params["page"],
			"PageContent": "This is " + params["page"] + " page.",
		})
	})

	m.RunOnAddr(":8080")
}
