package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, urlExist, methodExist := r.FindHandler(request.URL.Path, request.Method)

	if !urlExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)
}

func (r *Router) FindHandler(path, method string) (http.HandlerFunc, bool, bool) {
	_, pathExist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, pathExist, methodExist
}