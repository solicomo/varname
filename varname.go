package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"runtime/debug"
	"sync"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	Theme      string
	PortalType string
	PortalAddr string
}

type VarName struct {
	root    string
	appName string
	config  Config
	db      *sql.DB
}

func (self *VarName) Init(root, appName string) (err error) {

	self.root = root
	self.appName = appName

	err = self.initConfig()

	if err != nil {
		return err
	}

	err = self.initDB()

	return
}

func (self *VarName) initConfig() (err error) {

	configFile := path.Join(self.root, self.appName+".conf")

	configData, err := ioutil.ReadFile(configFile)

	if err != nil {
		return
	}

	err = json.Unmarshal(configData, &self.config)

	return
}

func (self *VarName) initDB() (err error) {

	self.db, err = sql.Open("sqlite3", path.Join(self.root, "app/data", self.appName+".db"))

	return
}

func (self *VarName) Run() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		self.runPortal()
	}()

	go func() {
		defer wg.Done()
		self.timeWork()
	}()

	wg.Wait()
}

func (self *VarName) timeWork() {

}

func check(err error) {
	if err != nil {
		log.Println("[ERRO]", err)
		log.Println("[DEBU]", string(debug.Stack()[:]))
		panic(err)
	}
}

type handlerFunc func(*http.Request, martini.Params, map[string]interface{})
type handlerWrapFunc func(*http.Request, martini.Params, render.Render)

func martiniSafeHandler(layout string, hf handlerFunc) handlerWrapFunc {

	return func(req *http.Request, params martini.Params, r render.Render) {

		data := make(map[string]interface{})

		defer func() {
			if err, ok := recover().(error); ok {
				data["Status"] = "500"
				data["ErrMsg"] = err.Error()

				r.HTML(500, layout, data)
			}
		}()

		log.Println("[INFO]", "request from", req.RemoteAddr, req.URL)

		hf(req, params, data)

		r.HTML(200, layout, data)
	}
}
