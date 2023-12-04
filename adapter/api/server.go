package api

import (
	"log"
	"myTestWithMultiStage/env"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	Run()
}

type server struct {
	port   string
	server *gin.Engine
}

func New() IServer {
	return &server{
		port:   env.Var.AppApiPort,
		server: gin.Default(),
	}
}

func (s *server) Run() {
	router := routes(s.server)

	log.Fatal(router.Run(":" + s.port))
}
