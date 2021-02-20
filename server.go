package main

import (
	"fmt"
	"go-echo-template/config"
	"go-echo-template/logger"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tylerb/graceful"
)

// Server context
type Server struct {
	Router *echo.Echo
	Cfg    *config.Config
}

// NewServer constructor
func NewServer() *Server {
	cfg := config.Get()

	return &Server{
		Router: GetRouter(),
		Cfg:    cfg,
	}
}

// Run server
func (s *Server) Run() {
	srv := s.Router
	srv.Server.Addr = fmt.Sprintf("%s:%s", s.Cfg.ServerHost, s.Cfg.ServerPort)

	logger.Log.Infof("Starting web server on port %s...", s.Cfg.ServerPort)
	graceful.ListenAndServe(srv.Server, 5*time.Second)
}
