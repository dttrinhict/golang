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

func NewGinServer(user *handler.User) *gin.Engine{
	engine := gin.New()
	internal := engine.Group("/")
	internal.POST("/users", user.GUserCreate)
	internal.GET("/users", user.GGetUsers)
	//internal.GET("/users", user.GetUsers)
	//internal.GET("/user/:id", user.GetUser)
	//internal.PUT("/user", user.UpdateUser)
	return engine
}