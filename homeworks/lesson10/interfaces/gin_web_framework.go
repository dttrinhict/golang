package interfaces

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang/homeworks/lesson10/interfaces/handler"
	"golang/homeworks/lesson10/util/logger"
	"log"
	"net/http"
)


type server struct {
	httpServer *http.Server
}

var Logger logger.Logger

func (s *server) Run() {
	//shutdown.SigtermHandler().RegisterErrorFuncContext(context.Background(), s.httpServer.Shutdown)
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Panic("Server listen and serve error", err)
	}
}

func NewGinServer(user *handler.User, club *handler.Club, userClub *handler.UserClub, member *handler.Member) *gin.Engine{
	engine := gin.New()
	userGroup := engine.Group("/user/")
	{
		userGroup.POST("/create", user.GUserCreate)
		userGroup.GET("/get-users", user.GGetUsers)
		userGroup.GET("/get-user/:id", user.GGetUser)
		userGroup.PUT("/update-user", user.GUpdateUser)
	}
	memberGroup := engine.Group("/member/")
	{
		memberGroup.POST("/create", member.GMemberCreate)
		memberGroup.GET("/get-members", member.GGetMembers)
		memberGroup.GET("/get-member/:id", member.GGetMember)
		memberGroup.PUT("/update-member", member.GUpdateMember)
	}
	clubGroup := engine.Group("/club/")
	{
		clubGroup.POST("/create", club.GClubCreate)
		clubGroup.GET("/get-clubs", club.GGetClubs)
		clubGroup.GET("/get-club/:id", club.GGetClub)
		clubGroup.PUT("/update-club", club.GUpdateClub)
	}
	userClubGroup := engine.Group("/user-club/")
	{
		userClubGroup.POST("/assign-user-to-club", userClub.GAssignUserToClub)
		userClubGroup.GET("/get-users-of-club/:id", userClub.GGetUsersOfClub)
	}

	return engine
}