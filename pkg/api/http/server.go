package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServerConfig struct {
	Address     string
	Environment string
}

type Server struct {
	ServerConfig
	engine     *gin.Engine
	httpServer *http.Server
}

func NewServer(cfg ServerConfig, configureRoutes func(engine *gin.Engine)) Server {
	engine := gin.New()
	engine.ContextWithFallback = true
	configureRoutes(engine)
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	httpServer := &http.Server{
		Addr:    cfg.Address,
		Handler: engine,
	}

	return Server{engine: engine, httpServer: httpServer, ServerConfig: cfg}
}

func (s Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
