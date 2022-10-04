package sap

import (
	"log"
	"net/http"
	"strings"
)

type Path struct {
	Url    string
	Get    func(c Ctx)
	Post   func(c Ctx)
	Put    func(c Ctx)
	Delete func(c Ctx)
}

func GenerateRoute(p Path) {
	router.HandleFunc(strings.Split(p.Url, "(")[0], func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method + " request made at " + r.URL.Path)
		switch r.Method {
		case "GET":
			p.Get(Ctx{Httpobj{w, r}, URLPARSE(p.Url, r.URL.Path)})
		case "POST":
			p.Post(Ctx{Httpobj{w, r}, URLPARSE(p.Url, r.URL.Path)})
		case "PUT":
			p.Put(Ctx{Httpobj{w, r}, URLPARSE(p.Url, r.URL.Path)})
		case "DELETE":
			p.Delete(Ctx{Httpobj{w, r}, URLPARSE(p.Url, r.URL.Path)})
		}
	})
}
