package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TechMaster/eris"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/v12"
	"log"
	"net/http"
	"syscall"
)


type Logger interface {
	Log(err error)
	Info(msg string, data interface{}, redirectLink ...string)
}

var Log Logger

type GinLogger struct {
	GinCtx gin.Context
}



func NewGinLogger(ginCtx *gin.Context) Logger  {
	if Log == nil {
		 return &GinLogger{
			GinCtx: *ginCtx,
		}
	}
	return Log
}


func (g GinLogger) Log(err error) {
	if errors.Is(err, syscall.EPIPE) {
		return
	}
	//Trả về JSON error khi client gọi lên bằng AJAX hoặc request.ContentType dạng application/json
	shouldReturnJSON := g.GinCtx.ContentType() == "application/json"
	switch e := err.(type) {
	case *eris.Error:
		if e.ErrType > eris.WARNING { //Chỉ log ra console hoặc file
			logErisError(e)
		}

		if shouldReturnJSON { //Có trả về báo lỗi dạng JSON cho REST API request không?
			errorBody := gin.H{
				"error": e.Error(),
			}
			if e.Data != nil { //không có dữ liệu đi kèm thì chỉ cần in thông báo lỗi
				errorBody["data"] = e.Data
			}
			if e.Code > 300 {
				g.GinCtx.Status(e.Code)
			} else {
				g.GinCtx.Status(http.StatusInternalServerError)
			}

			g.GinCtx.JSON(e.Code, errorBody) //Trả về cho client gọi REST API
			return                     //Xuất ra JSON rồi thì không hiển thị Error Page nữa
		}

		// Nếu request không phải là REST request (AJAX request) thì render error page
		//g.GinCtx.ViewData("ErrorMsg", e.Error())
		//if e.Data != nil {
		//	if bytes, err := json.Marshal(e.Data); err == nil {
		//		i.irisCtx.ViewData("Data", string(bytes))
		//	}
		//}
		//_ = i.irisCtx.View(LogConf.ErrorTemplate)
		//return
	default: //Lỗi thông thường
		fmt.Println(err.Error()) //In ra console
		if shouldReturnJSON {    //Trả về JSON
			//g.GinCtx.Status(http.StatusInternalServerError)
			g.GinCtx.JSON(http.StatusInternalServerError, err.Error())
		} else {
			g.GinCtx.JSON(http.StatusInternalServerError, gin.H{
					"ErrorMsg": err.Error(),
				})
		}
		return
	}
}

func (g GinLogger) Info(msg string, data interface{} , redirectLink ...string) {
	log.Printf("%v", map[string]interface{}{
		"Client": g.GinCtx.ClientIP(),
		"Message": msg,
		"Data": data,
		"Uri": redirectLink,
	})
}
//func (g GinLogger) Log(err error) {
//	var ctx gin.Context
//	panic("implement me")
//	if errors.Is(err, syscall.EPIPE) {
//		return
//	}
//}



type IrisLogger struct {
	irisCtx iris.Context
}


func NewIrisLogger(irisCtx *iris.Context) Logger  {
	if Log == nil {
		return &IrisLogger{
			irisCtx: *irisCtx,
		}
	}
	return Log
}

func (i IrisLogger) Log(err error) {
	if errors.Is(err, syscall.EPIPE) {
		return
	}
	//Trả về JSON error khi client gọi lên bằng AJAX hoặc request.ContentType dạng application/json
	shouldReturnJSON := i.irisCtx.IsAjax() || i.irisCtx.GetContentTypeRequested() == "application/json"
	switch e := err.(type) {
	case *eris.Error:
		if e.ErrType > eris.WARNING { //Chỉ log ra console hoặc file
			logErisError(e)
		}
		if shouldReturnJSON { //Có trả về báo lỗi dạng JSON cho REST API request không?
			errorBody := iris.Map{
				"error": e.Error(),
			}
			if e.Data != nil { //không có dữ liệu đi kèm thì chỉ cần in thông báo lỗi
				errorBody["data"] = e.Data
			}
			if e.Code > 300 {
				i.irisCtx.StatusCode(e.Code)
			} else {
				i.irisCtx.StatusCode(iris.StatusInternalServerError)
			}

			_, _ = i.irisCtx.JSON(errorBody) //Trả về cho client gọi REST API
			return                     //Xuất ra JSON rồi thì không hiển thị Error Page nữa
		}

		// Nếu request không phải là REST request (AJAX request) thì render error page
		i.irisCtx.ViewData("ErrorMsg", e.Error())
		if e.Data != nil {
			if bytes, err := json.Marshal(e.Data); err == nil {
				i.irisCtx.ViewData("Data", string(bytes))
			}
		}
		_ = i.irisCtx.View(LogConf.ErrorTemplate)
		return
	default: //Lỗi thông thường
		fmt.Println(err.Error()) //In ra console
		if shouldReturnJSON {    //Trả về JSON
			i.irisCtx.StatusCode(iris.StatusInternalServerError)
			_, _ = i.irisCtx.JSON(err.Error())
		} else {
			_ = i.irisCtx.View(LogConf.ErrorTemplate, iris.Map{
				"ErrorMsg": err.Error(),
			})
		}
		return
	}
}

func (i IrisLogger) Info(msg string, data interface{}, redirectLink ...string) {
	switch len(redirectLink) {
	case 2:
		_ = i.irisCtx.View(LogConf.InfoTemplate, iris.Map{
			"Msg":       msg,
			"LinkTitle": redirectLink[1], //<a href='Link'>LinkTitle</a>
			"Link":      redirectLink[0],
		})
	case 1:
		_ = i.irisCtx.View(LogConf.InfoTemplate, iris.Map{
			"Msg":       msg,
			"LinkTitle": redirectLink[0], //<a href='Link'>LinkTitle</a>
			"Link":      redirectLink[0],
		})
	default:
		_ = i.irisCtx.View(LogConf.InfoTemplate, iris.Map{
			"Msg": msg,
		})
	}
}