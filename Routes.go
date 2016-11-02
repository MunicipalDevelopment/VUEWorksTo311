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
		"IncomingVUEWorksServiceReqClosed",
		"POST",
		"/",
		IncomingVUEWorksServiceReqClosed,
	},
	Route{
		"isAlive",
		"GET",
		"/isAlive",
		isAlive,
	},
	Route{
		"total",
		"GET",
		"/total",
		total,
	},
}
