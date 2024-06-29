package httpserver

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

type Server struct {
	srv  *http.Server
	conf Config
}

func New(conf Config, handler http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Handler:      handler,
			ReadTimeout:  time.Duration(conf.ReadTimeoutMs) * time.Millisecond,
			WriteTimeout: time.Duration(conf.WriteTimeoutMs) * time.Millisecond,
			Addr:         fmt.Sprintf(":%v", conf.Port),
		},
		conf: conf,
	}
}

func (s *Server) Start() error {
	go func() {
		log.Printf("Starting http server on port %d", s.conf.Port)

		err := s.srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Failed to start http server due to: %v", err)
		}
	}()

	return nil
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.GracefulShutdownTimeoutMs)*time.Millisecond)
	defer cancel()

	log.Print("Shutting down http server")

	return s.srv.Shutdown(ctx)
}
