package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hmaier-dev/http-tool/internal/handlers"

	// blank import for handlers. They initalize theirself by init()
	_ "github.com/hmaier-dev/http-tool/internal/handlers/user"
)


type Server struct {
	Router *mux.Router
}

func NewServer() *Server {
	router := mux.NewRouter()

	for _, h := range handlers.GetHandlers() {
		// Link the routes declared in sub-handlers to *mux.Router
		h.Routes(router)
	}
  router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	// logs all routes when starting after they go defined
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		method, _ := route.GetMethods()
		log.Println(method, path)
		return nil
	})

	return &Server{Router: router}
}
