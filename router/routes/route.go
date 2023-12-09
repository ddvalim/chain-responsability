package routes

import "net/http"

type Route struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
}
