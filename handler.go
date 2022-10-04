package sap

import "net/http"

/*
Handle incoming Get requests at a given url
*/
func (s Sap) Get(url string, handler func(c Ctx)) {
	if len(routes) == 0 {
		routes = append(routes, Path{url, handler, nil, nil, nil})
	} else {
		for i, v := range routes {
			if v.Url == url {
				routes[i].Get = handler
				break
			} else if i == len(routes)-1 {
				routes = append(routes, Path{url, handler, nil, nil, nil})
				break
			}
		}
	}
}

/*
Handle incoming Post requests at a given url
*/
func (s Sap) Post(url string, handler func(c Ctx)) {
	if len(routes) == 0 {
		routes = append(routes, Path{url, nil, handler, nil, nil})
	} else {
		for i, v := range routes {
			if v.Url == url {
				routes[i].Post = handler
				break
			} else if i == len(routes)-1 {
				routes = append(routes, Path{url, nil, handler, nil, nil})
				break
			}
		}
	}
}

/*
Handle incoming Put requests at a given url
*/
func (s Sap) Put(url string, handler func(c Ctx)) {
	if len(routes) == 0 {
		routes = append(routes, Path{url, nil, nil, handler, nil})
	} else {
		for i, v := range routes {
			if v.Url == url {
				routes[i].Put = handler
				break
			} else if i == len(routes)-1 {
				routes = append(routes, Path{url, nil, nil, handler, nil})
				break
			}
		}
	}
}

/*
Handle incoming Delete requests at a given url
*/
func (s Sap) Delete(url string, handler func(c Ctx)) {
	if len(routes) == 0 {
		routes = append(routes, Path{url, nil, nil, nil, handler})
	} else {
		for i, v := range routes {
			if v.Url == url {
				routes[i].Delete = handler
				break
			} else if i == len(routes)-1 {
				routes = append(routes, Path{url, nil, nil, nil, handler})
				break
			}
		}
	}
}

/*
Create a fileserver to handle static file serving
*/
func (s Sap) FileServer(url string, path string) {
	var fs = http.FileServer(http.Dir(path))
	router.Handle(url+"/", http.StripPrefix(url, fs))
}
