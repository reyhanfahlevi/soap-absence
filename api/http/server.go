package http

import (
	"context"
	"net/http"

	"time"

	"github.com/reyhanfahlevi/soap-absence/api"
	"github.com/reyhanfahlevi/soap-absence/api/http/absence"
	"github.com/tokopedia/affiliate/pkg/listener"
)

// Server struct
type Server struct {
	server     *http.Server
	AbsenceSvc api.AbsenceService
}

// Serve will run an HTTP server
func (s *Server) Serve(port string) error {

	// init absence service
	absence.Init(s.AbsenceSvc)

	s.server = &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      handler(),
	}

	lis, err := listener.Listen(port)
	if err != nil {
		return err
	}

	return s.server.Serve(lis)
}

// Shutdown will tear down the server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
