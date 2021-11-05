package logger

import (
	"errors"
	"fmt"
	"github.com/TechMaster/eris"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"syscall"
)

type FiberLogger struct {
	FiberCtx fiber.Ctx
}



func NewFiberLogger(fiberCtx *fiber.Ctx) Logger  {
	if Log == nil {
		return &FiberLogger{
			FiberCtx: *fiberCtx,
		}
	}
	return Log
}


func (f FiberLogger) Log(err error) {
	if errors.Is(err, syscall.EPIPE) {
		return
	}
	switch e := err.(type) {
	case *eris.Error:
		if e.ErrType > eris.WARNING { //Chỉ log ra console hoặc file
			logErisError(e)
		}

		errorBody := fiber.Map{
			"error": e.Error(),
		}
		if e.Data != nil { //không có dữ liệu đi kèm thì chỉ cần in thông báo lỗi
			errorBody["data"] = e.Data
		}
		if e.Code > 300 {
			f.FiberCtx.Status(e.Code)
		} else {
			f.FiberCtx.Status(http.StatusInternalServerError)
		}

		f.FiberCtx.JSON(errorBody) //Trả về cho client gọi REST API
		return                     //Xuất ra JSON rồi thì không hiển thị Error Page nữa
	default: //Lỗi thông thường
		fmt.Println(err.Error()) //In ra console
		f.FiberCtx.Status(http.StatusInternalServerError)
		f.FiberCtx.JSON(fiber.Map{
			"ErrorMsg": err.Error(),
		})
		return
	}
}

func (f FiberLogger) Info(msg string, data interface{}, redirectLink ...string) {
	log.Printf("%v", map[string]interface{}{
		"Client": f.FiberCtx,
		"Message": msg,
		"Data": data,
		"Uri": redirectLink,
	})
}