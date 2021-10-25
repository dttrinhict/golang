package interfaces

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang/homeworks/lesson10/interfaces/handler"
	"log"
	"net/http"
)


type server struct {
	httpServer *http.Server
}

func (s *server) Run() {
	//shutdown.SigtermHandler().RegisterErrorFuncContext(context.Background(), s.httpServer.Shutdown)
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Panic("Server listen and serve error", err)
	}
}

func NewGinServer(user *handler.User, club *handler.Club) *gin.Engine{
	engine := gin.New()
	userGroup := engine.Group("/user/")
	{
		userGroup.POST("/create", user.GUserCreate)
		userGroup.GET("/get-users", user.GGetUsers)
		userGroup.GET("/get-user/:id", user.GGetUser)
		userGroup.PUT("/update-user", user.GUpdateUser)
	}
	clubGroup := engine.Group("/club/")
	{
		clubGroup.POST("/create", club.GClubCreate)
		clubGroup.GET("/get-clubs", club.GGetClubs)
		clubGroup.GET("/get-club/:id", club.GGetClub)
		clubGroup.PUT("/update-club", club.GUpdateClub)
	}

	return engine
}