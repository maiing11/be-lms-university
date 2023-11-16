package delivery

import (
	"fmt"
	"log"

	"enigmacamp.com/be-lms-university/config"
	"enigmacamp.com/be-lms-university/delivery/controller"
	"enigmacamp.com/be-lms-university/manager"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *gin.Engine
	host      string
}

func (s *Server) setupController() {
	rg := s.engine.Group("/api/v1")
	// semua controller didaftarkan disini
	controller.NewEnrollmentController(s.ucManager.EnrollmentUseCase(), rg).Route()
}

func (s *Server) Run() {
	s.setupController()
	if err := s.engine.Run(s.host); err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	// Infra Manager
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	infraManager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		ucManager: useCaseManager,
		engine:    engine,
		host:      host,
	}
}
