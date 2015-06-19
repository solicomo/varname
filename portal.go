package main

import (
	"net/http"
	"path"
	"varname/app/controllers"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func (self *VarName) runPortal() {

	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory:  path.Join(self.root, "app/views/"+self.config.Theme),
		Extensions: []string{".gohtml", ".tmpl", ".html"},
	}))

	m.Get("/search", martiniSafeHandler("search", self.handleSearch))
	m.Get("/e/:entry", martiniSafeHandler("entry", self.handleEntry))
	m.Get("/(?P<page>(about|terms))", martiniSafeHandler("page", self.handlePage))

	m.Get("/**", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/index.html", http.StatusTemporaryRedirect)
	})

	m.RunOnAddr(self.config.PortalAddr)
}

func (self *VarName) handleIndex(req *http.Request, params martini.Params, data map[string]interface{}) {

}

func (self *VarName) handleSearch(req *http.Request, params martini.Params, data map[string]interface{}) {

	var err error

	qs := req.FormValue("q")
	qs = controllers.Normalize(qs)

	data["Title"] = qs
	data["PageClass"] = "search"
	data["SearchString"] = qs
	data["SearchRelates"] = []string{qs, qs, qs, qs}
	data["SearchList"], err = controllers.Search(0, 10, qs)

	check(err)
}

func (self *VarName) handleEntry(req *http.Request, params martini.Params, data map[string]interface{}) {

	data["Entry"] = params["entry"]
}

func (self *VarName) handlePage(req *http.Request, params martini.Params, data map[string]interface{}) {

	data["PageHead"] = params["page"]
	data["PageContent"] = "This is " + params["page"] + " page."
}
