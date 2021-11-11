package logger

import (
	"os"

	"github.com/TechMaster/eris"
)

type LogConfig struct {
	LogFolder     string // thư mục chứa log file. Nếu rỗng có nghĩa là không ghi log ra file
	ErrorTemplate string // tên view template sẽ render error page
	InfoTemplate  string // tên view template sẽ render info page
	Top           int    // số dòng đỉnh stack trace sẽ được in ra
}

var LogConf *LogConfig
var logFile *os.File

var ErisStringFormat eris.StringFormat

//Dùng variadic param chỉ để tạo ra một optional param: có thể có hoặc không
//Nếu không truyền tham số thì sẽ tự tạo tham số mặc định
func Init(logConfig ...LogConfig) *os.File {
	if len(logConfig) > 0 {
		LogConf = &logConfig[0]
	} else { //Truyền cấu hình nil thì tạo cấu hình mặc định
		LogConf = &LogConfig{
			LogFolder:     "logs/", // thư mục chứa log file. Nếu rỗng có nghĩa là không ghi log ra file
			ErrorTemplate: "error", // tên view template sẽ render error page
			InfoTemplate:  "info",  // tên view template sẽ render info page
			Top:           3,       // số dòng đầu tiên trong stack trace sẽ được giữ lại
		}
	}

	ErisStringFormat = eris.StringFormat{
		Options: eris.FormatOptions{
			InvertOutput: false, // flag that inverts the error output (wrap errors shown first)
			WithTrace:    true,  // flag that enables stack trace output
			InvertTrace:  true,  // flag that inverts the stack trace output (top of call stack shown first)
			WithExternal: false,
			Top:          LogConf.Top, // Chỉ lấy 3 dòng lệnh đầu tiên
			//Mục tiêu để báo lỗi gọn hơn, stack trace đủ ngắn
		},
		MsgStackSep:  "\n",  // separator between error messages and stack frame data
		PreStackSep:  "\t",  // separator at the beginning of each stack frame
		StackElemSep: " | ", // separator between elements of each stack frame
		ErrorSep:     "\n",  // separator between each error in the chain
	}

	if LogConf.LogFolder != "" {
		deleteZeroByteLogFiles(LogConf.LogFolder)
		logFile = newLogFile(LogConf.LogFolder)
		return logFile
	} else {
		return nil
	}
}
