package controllers

import (
	"log"
	"strings"
	"varname/app/models"
)

type SearchResult struct {
	Title   string
	Url     string
	Caption string
}

type SearchNav struct {
	Name  string
	Class string
	Start int
}

func Normalize(qs string) (ns string) {

	ns = strings.Join(strings.Fields(qs), " ")
	return
}

func Search(offset, count int, qs string) (results []SearchResult, navs []SearchNav, err error) {

	var entries models.Entries

	es, total, err := entries.Search(offset, count, qs)

	if err != nil {
		log.Println("[ERRO]", err)
		return
	}

	results = make([]SearchResult, 0)

	for _, e := range es {

		var res SearchResult

		res.Title = e.Name
		res.Url = e.Slug
		res.Caption = e.Caption

		results = append(results, res)
	}

	navs = make([]SearchNav, 0)

	return
}
