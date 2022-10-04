package sap

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
