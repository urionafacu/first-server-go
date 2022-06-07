package main

import (
	"net/http"
)

type Server struct {
	port string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
		router: NewRouter(),
	}
}

func (s *Server) AddMidleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	return http.ListenAndServe(s.port, nil)
}

func (s *Server) Handle(method, path string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path]

	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}

	s.router.rules[path][method] = handler
}