package sap

import (
	"fmt"
	"log"
	"net/http"
)

type Sap struct{}

/*
Create a new instance of a sap router via this function
*/
func Init() Sap {
	return Sap{}
}

/*
Run your app on a given port using this function
*/
func (s Sap) Run(port int) {
	for _, v := range routes {
		GenerateRoute(v)
	}
	log.Println("Starting server at port " + fmt.Sprint(port) + "...")
	http.ListenAndServe("localhost:"+fmt.Sprint(port), router)
}
