package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

/*
Loại bỏ những log file cũ nhưng không có dữ liệu được ghi toàn là zero byte
*/
func deleteZeroByteLogFiles(logFolder string) {
	files, err := ioutil.ReadDir(logFolder)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && file.Size() == 0 {
			os.Remove(logFolder + "/" + file.Name())
		}
	}
}
func todayFilename() string {
	today := time.Now().Format("2006 01 02")
	return today + ".txt"
}

func newLogFile(logFolder string) *os.File {
	if logFolder == "" {
		return nil
	}
	filename := todayFilename()
	// Open the file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(logFolder+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}