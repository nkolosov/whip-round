package server

import (
	"context"
	"database/sql"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	db         *sql.DB
}

func NewServer(cfg *http.Server) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           cfg.Addr,
			Handler:        cfg.Handler,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: cfg.MaxHeaderBytes,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	err := s.CloseDBConnection()
	if err != nil {
		return err
	}

	return s.httpServer.Shutdown(ctx)
}

func (s *Server) CloseDBConnection() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
