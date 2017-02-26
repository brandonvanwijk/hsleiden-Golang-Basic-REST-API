package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"find",
		"GET",
		"/todos",
		find,
	},
	Route{
		"findOne",
		"GET",
		"/todos/{todoId}",
		findOne,
	},
	Route{
		"create",
		"POST",
		"/todos",
		create,
	},
}
