package server

import (
	"fmt"
	"net/http"
	"strings"
)

// Server struct
type Server struct {
	router *Router
}

// NewServer returns new server
func NewServer() (*Server, error) {

	s := &Server{
		router: NewRouter(),
	}

	return s, s.init()
}

// Innitialize the server and add routing
func (s *Server) init() error {

	// // Add route for main page
	// s.router.Route("/")["GET"] = http.HandlerFunc(s.handleIndexGET)

	// // Getter pages
	// s.router.Route("/users")["GET"] = http.HandlerFunc(s.handleUsersGET)
	// s.router.Route("/users/ws")["GET"] = http.HandlerFunc(s.handleUserWS)

	// s.router.Route("/comments")["GET"] = http.HandlerFunc(s.handleCommentsGET)
	// s.router.Route("/comments/ws")["GET"] = http.HandlerFunc(s.handleCommentsWS)
	return nil
}

// ServeHTTP to the server
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	if !prefixChecker(url, "/style", "/favicon.ico") {
		p := Path(r.URL.Path)
		m := Method(r.Method)

		fmt.Println(p, m)

		s.router.Route(p)[m].ServeHTTP(w, r)
		return
	}

	s.handleGETAssets(w, r)
}

// If assets are called -> host assets
func (s Server) handleGETAssets(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("../server/assets/")).ServeHTTP(w, r)
}

// prefixChecker checks if any of given prefixes is in the url
func prefixChecker(url string, prefix ...string) bool {
	out := false
	for _, p := range prefix {
		if strings.HasPrefix(url, p) {
			out = true
		}
	}
	return out
}
