package delivery

import (
	"database/sql"
	"fmt"

	"enigmacamp.com/be-lms-university/delivery/controller"
	"enigmacamp.com/be-lms-university/repository"
	"enigmacamp.com/be-lms-university/usecase"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	uc     usecase.EnrollmentUseCase
	engine *gin.Engine
	host   string
}

func (s *Server) setupController() {
	rg := s.engine.Group("/api/v1")
	// semua controller didaftarkan disini
	controller.NewEnrollmentController(s.uc, rg).Route()
}

func (s *Server) Run() {
	s.setupController()
	if err := s.engine.Run(s.host); err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbName := "lms_university"
	driver := "postgres"

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	// enrollment
	enrollRepo := repository.NewEnrollmentRepository(db)
	userRepo := repository.NewUserRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	userUC := usecase.NewUserUseCase(userRepo)
	courseUC := usecase.NewCourseUseCase(courseRepo)
	enrollmentUC := usecase.NewEnrollmentUseCase(enrollRepo, userUC, courseUC)

	engine := gin.Default()
	// ini kalo misalkan portnya mau diubah
	apiHost := fmt.Sprintf(":%s", "8080")

	return &Server{
		uc:     enrollmentUC,
		engine: engine,
		host:   apiHost,
	}
}
