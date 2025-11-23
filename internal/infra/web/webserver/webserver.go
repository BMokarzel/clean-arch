package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      *Handlers
	WebServerPort string
}

type Handlers struct {
	List []Handler
}

type Handler struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      &Handlers{},
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	s.Handlers.List = append(s.Handlers.List, Handler{
		Method:  method,
		Path:    path,
		Handler: handler,
	})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers.List {
		switch handler.Method {
		case http.MethodGet:
			s.Router.Get(handler.Path, handler.Handler)
		case http.MethodPut:
			s.Router.Put(handler.Path, handler.Handler)
		case http.MethodPost:
			s.Router.Post(handler.Path, handler.Handler)
		case http.MethodPatch:
			s.Router.Patch(handler.Path, handler.Handler)
		case http.MethodDelete:
			s.Router.Delete(handler.Path, handler.Handler)
		default:
			continue
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
