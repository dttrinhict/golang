package logger
import (
	"fmt"
	"time"

	"github.com/TechMaster/eris"

	"github.com/goccy/go-json"
)

//Hàm chuyên xử lý Eris Error có Stack Trace. Chỉ áp dụng với cấp độ lỗi ERROR, SYSERROR, PANIC
func logErisError(err *eris.Error) {
	formattedStr := eris.ToCustomString(err, ErisStringFormat) //Định dạng lỗi Eris

	//Chỗ này log ra console
	if err.ErrType > eris.ERROR { //Với lỗi cao hơn ERROR gồm SYSERROR và PANIC, in ra mầu đỏ và ghi ra file
		colorReset := string("\033[0m")
		colorMagenta := string("\033[35m")
		fmt.Println(colorMagenta, formattedStr, colorReset)

		dataString := marshalErisData2JSON(err)

		if dataString != "" { //Nếu có dữ liệu đi kèm thì cũng ghi ra file
			fmt.Println(colorMagenta, dataString, colorReset)
		}

		//Lỗi Panic và Error nhất thiết phải ghi vào file. Và chỉ ghi khi LogFolder được cài đặt
		if logFile != nil {
			var textToFile string
			if dataString != "" { //Nếu có dữ liệu đi kèm thì cũng ghi ra file
				textToFile = time.Now().Format("2006 01 02-15:04:05 - ") + formattedStr + "\n" + dataString + "\n\n"
			} else {
				textToFile = time.Now().Format("2006 01 02-15:04:05 - ") + formattedStr + "\n\n"
			}
			if _, err := logFile.WriteString(textToFile); err != nil {
				panic(err)
			}
		}

	} else {
		fmt.Println(formattedStr) //Error Level
	}
}

func marshalErisData2JSON(err *eris.Error) string {
	if err.Data != nil {
		if dataStr, err := json.Marshal(err.Data); err == nil {
			return string(dataStr)
		}
	}
	return ""
}
