package server

import (
	"net/http"
	"time"

	"github.com/tomaszkoziara/mvc-go/handlers"

	"github.com/gorilla/mux"
)

// Server is the server.
type Server struct {
	userHandler      handlers.IUser
	dashboardHandler handlers.IDashboard
}

// Serve serves all routes.
func (s *Server) Serve() error {
	r := mux.NewRouter()

	r.HandleFunc("/dashboard", s.dashboardHandler.Show).Methods("GET")

	apiPrefix := r.PathPrefix("/api").Subrouter()
	apiPrefix.HandleFunc("/user", s.userHandler.Create).Methods("POST")
	apiPrefix.HandleFunc("/users", s.userHandler.List).Methods("GET")

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	return srv.ListenAndServe()
}

// New creates a new Server.
func New(
	userHandler handlers.IUser,
	dashboardHandler handlers.IDashboard,
) *Server {

	return &Server{
		userHandler:      userHandler,
		dashboardHandler: dashboardHandler,
	}
}
