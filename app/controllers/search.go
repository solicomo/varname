package controllers

import "strings"

type SearchResult struct {
	Title   string
	Url     string
	Caption string
}

func Normalize(qs string) (ns string) {

	ns = strings.Join(strings.Fields(qs), " ")
	return
}

func Search(offset, count int, qs string) (results []SearchResult, err error) {

	return
}
