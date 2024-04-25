package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prathameshj.dev/passhash/db"
	"prathameshj.dev/passhash/models"
)

type Server interface {
	Start()
	Readiness(ctx *gin.Context)
	Liveness(ctx *gin.Context)

	GetAllWebsites(ctx *gin.Context)
	GetPasswordByWebsite(ctx *gin.Context)
	GeneratePassword(ctx *gin.Context)
	AddPassword(ctx *gin.Context)
	DeletePassword(ctx *gin.Context)
}

type GinServer struct {
	gin *gin.Engine
	DB  db.DatabaseClient
}

func StartServer(db db.DatabaseClient) Server {
	server := &GinServer{
		gin: gin.Default(),
		DB:  db,
	}
	return server
}

func (s *GinServer) Start() {
	s.registerRoutes()
	s.gin.Run(":8080")
}

func (s *GinServer) registerRoutes() {
	s.gin.GET("/readiness", s.Readiness)
	s.gin.GET("/liveness", s.Liveness)

	s.gin.GET("/getWebsites", s.GetAllWebsites)
	s.gin.GET("/getPasswordByWebsite", s.GetPasswordByWebsite)
	s.gin.POST("/generateNewPassword", s.GeneratePassword)
	s.gin.POST("/addPassword", s.AddPassword)
	s.gin.DELETE("/deletePassword", s.DeletePassword)

}

func (s *GinServer) Readiness(ctx *gin.Context) {
	ready := s.DB.Ready()
	if ready {
		ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
		return
	}

	ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure"})
}

func (s *GinServer) Liveness(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}
