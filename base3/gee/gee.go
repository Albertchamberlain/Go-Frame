package main

import (
	"fmt"
	"net/http"
)

//handle http func
type HandleFunc func(w http.ResponseWriter, r *http.Request)

// kv save map
type Engine struct {
	router map[string]HandleFunc
}

//New constructor of Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

//get define the method to add Get request
func (engine *Engine) GET(pattern string, handler HandleFunc) {
	engine.addRoute("GET", pattern, handler)
}

//post define the method to add post request
func (engine *Engine) POST(pattern string, handler HandleFunc) {
	engine.addRoute("POST", pattern, handler)
}

//defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

//impl ServeHTTP
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		_, _ = fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}
