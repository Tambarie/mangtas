package router

import (
	"fmt"
	"net/http"
)

type Router struct {
	ContentType  string
	handlers map[string]func(w http.ResponseWriter, r *http.Request)
}

func NewRouter() *Router {
	router := new(Router)
	router.handlers = make(map[string]func(w http.ResponseWriter, r *http.Request ))
	return router
}

func (s *Router) SetContentType (str string){
	s.ContentType = str
}

func (s *Router) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	f, ok := s.handlers[key(r.Method, r.URL.Path)]
	if !ok {
		notfound(w)
		return
	}

	w.Header().Set("Content-Type", s.ContentType)
	f(w,r)
}

func (s *Router) GET(path string, f http.HandlerFunc)  {
	s.handlers[key("GET",path)] = f
}



// POST handles post requests
func (s *Router) POST(path string, f http.HandlerFunc) {
	s.handlers[key("POST", path)] = f
}

// PUT handles put requests
func (s *Router) PUT(path string, f http.HandlerFunc) {
	s.handlers[key("PUT", path)] = f
}

// DELETE handles delete requests
func (s *Router) DELETE(path string, f http.HandlerFunc) {
	s.handlers[key("DELETE", path)] = f
}

func key(method, path string) string {
	return fmt.Sprintf("%s:%s", method, path)
}

func notfound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error":"not found"}`))
}

