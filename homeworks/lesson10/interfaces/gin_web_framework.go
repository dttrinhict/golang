package interfaces

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"golang/homeworks/lesson10/interfaces/handler"
	"golang/homeworks/lesson10/util/logger"
	"golang/homeworks/lesson10/util/storages"
	"io"
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

func NewGinServer(user *handler.User, club *handler.Club, userRole *handler.UserRole, member *handler.Member) *gin.Engine{
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
	userClubGroup := engine.Group("/user-role/")
	{
		userClubGroup.POST("/assign-user-to-role", userRole.GAssignUserToRole)
		userClubGroup.GET("/get-users-of-role/:id", userRole.GGetUsersOfRole)
	}
	download := engine.Group("/download/")
	{
		download.GET("/", func(c *gin.Context) {
			projectID := "project-ID"
			gcs, err := storages.GCSInit("service_account_key", projectID)
			if err!=nil {
				c.JSON(503, gin.H{
					"message": "could not access gcs",
				})
				return
			}
			object := "file_example_MP4_1920_18MG.mp4"
			bucketName := "trinhdt2-test"
			data, err := gcs.ReadFileFromBucket(object, bucketName)
			if err!=nil {
				c.JSON(503, gin.H{
					"message": "could not access gcs",
				})
				return
			}
			c.Header("Content-Description", "File Transfer")
			c.Header("Content-Transfer-Encoding", "binary")
			c.Header("Content-Disposition", "attachment; filename="+object )
			c.Header("Content-Type", "application/octet-stream")
			//c.Writer.Header().Add("Content-type", "application/octet-stream")
			reader := bytes.NewReader(data)
			_, err = io.Copy(c.Writer, reader)
			if err != nil {
				c.JSON(http.StatusNotFound,
					gin.H{
						"Message": err.Error(),
					})
				}
				return
			})
	}

	return engine
}