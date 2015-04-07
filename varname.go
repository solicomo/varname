package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory: "app/views/simple",
		Extensions: []string{".tmpl", ".html"},
	}))

//	m.Get("/", func(r render.Render) {
//		r.HTML(200, "index", "Variable Naming Service")
//	})

	m.Get("/:name", func(r render.Render, params martini.Params) {
		r.JSON(200, map[string]interface{}{"hello": params["name"]})
	})

	m.RunOnAddr(":8080")
}
