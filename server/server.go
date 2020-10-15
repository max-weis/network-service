package server

import (
	"github.com/julienschmidt/httprouter"
	"gitlab.com/M4xwell/network-service/host"
	"gitlab.com/M4xwell/network-service/pkg/logger"
	"net/http"
)

type Server struct {
	Logger logger.Logger
	router httprouter.Router
}

func NewServer(logger logger.Logger, service host.Service) Server {
	hostHandler := HostHandler{
		Logger:  logger,
		Service: service,
	}

	router := httprouter.New()
	router.GET("/host", hostHandler.ListHosts)
	router.GET("/host/:name", hostHandler.DetailHost)

	s := Server{
		Logger: logger,
		router: *router,
	}

	return s
}

func (s *Server) Serve() {
	s.Logger.Error(http.ListenAndServe(":8080", &s.router))

}
