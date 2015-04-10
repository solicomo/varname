package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

type SearchRecord struct {
	Title string
	Url string
	Caption string
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory: "app/views/simple",
		Extensions: []string{".tmpl", ".html"},
	}))

	m.Get("/search", func(req *http.Request, r render.Render) {
		qs := req.FormValue("q")
		data := make(map[string] interface {})
		data["Title"] = qs
		data["PageClass"] = "search"
		data["SearchString"] = qs
		data["SearchRelates"] = []string {qs, qs, qs, qs}
		data["SearchList"] = []SearchRecord {
			{qs, "/search?q=" + qs, qs + qs + qs},
			{qs, "/search?q=" + qs, qs + qs + qs},
			{qs, "/search?q=" + qs, qs + qs + qs},
			{qs, "/search?q=" + qs, qs + qs + qs},
			{qs, "/search?q=" + qs, qs + qs + qs},
			{qs, "/search?q=" + qs, qs + qs + qs},
			{qs, "/search?q=" + qs, qs + qs + qs},
		}

		r.HTML(200, "search", data)
	})

	m.Get("/:name", func(r render.Render, params martini.Params) {
		r.JSON(200, map[string]interface{}{"hello": params["name"]})
	})

	m.RunOnAddr(":8080")
}
