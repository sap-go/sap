package sap

import (
	"encoding/json"
	"net/http"
	"net/url"
	"text/template"
)

type Ctx struct {
	httpobj Httpobj
	Params  map[string]string
}

type Httpobj struct {
	w http.ResponseWriter
	r *http.Request
}

/*
Send a raw string response to the browser
*/
func (c Ctx) String(text string) {
	c.httpobj.w.Write([]byte(text))
}

/*
Send a json response to the browser
*/
func (c Ctx) Json(value any) error {
	var data, err = json.Marshal(value)
	c.httpobj.w.Header().Set("Content-Type", "application/json")
	c.httpobj.w.Write([]byte(data))
	return err
}

/*
Send an html string to the browser
*/
func (c Ctx) HTML(text string) {
	c.httpobj.w.Header().Set("Content-Type", "text/html")
	c.httpobj.w.Write([]byte(text))
}

/*
Send an html file to the browser
*/
func (c Ctx) HTMLFile(path string, data any) error {
	c.httpobj.w.Header().Set("Content-Type", "text/html")
	var render, err = template.ParseFiles(path)
	render.Execute(c.httpobj.w, data)
	return err
}

/*
Obtain the post body using this function
*/
func (c Ctx) Body(value any) {
	var decoder = json.NewDecoder(c.httpobj.r.Body)
	decoder.Decode(value)
}

/*
Obtain a given query param of a key using this function
*/
func (c Ctx) Query(key string) string {
	return c.httpobj.r.URL.Query()[key][0]
}

/*
Obtain all the query params using this function
*/
func (c Ctx) QueryAll() url.Values {
	return c.httpobj.r.URL.Query()
}

/*
Redirect the user from the given url to another one
*/
func (c Ctx) Redirect(url string) {
	http.Redirect(c.httpobj.w, c.httpobj.r, url, http.StatusMovedPermanently)
}

/*
Write the header code of a given url using this function
*/
func (c Ctx) WriteHeader(code int) {
	c.httpobj.w.WriteHeader(code)
}

/*
Set the header key/value of a given url using this function
*/
func (c Ctx) SetHeader(key string, value string) {
	c.httpobj.w.Header().Set(key, value)
}
