package server

import (
	"context"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(handler http.Handler, port string) Server {
	return Server{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: handler,
		},
	}
}

func (s Server) Run() error { return s.server.ListenAndServe() }

func (s Server) Shutdown(ctx context.Context) error { return s.server.Shutdown(ctx) }
