package sap

import (
	"strings"
)

func URLPARSE(url string, rawurl string) map[string]string {
	if !(strings.Contains(url, "(")) {
		return map[string]string{}
	} else {
		var params = make(map[string]string, 100)
		var urlarr = strings.Split(url, "(")[1:]
		var rawurlarr = strings.Split(strings.Split(rawurl, strings.Split(url, "(")[0])[1], "/")
		for i, v := range urlarr {
			params[strings.Split(v, ")")[0]] = rawurlarr[i]
		}
		return params
	}
}
