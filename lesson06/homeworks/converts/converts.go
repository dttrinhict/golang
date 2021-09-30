package converts

import (
	"fmt"
	"github.com/dttrinhict/golang/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	//"path/filepath"
	//"strings"
)

var entriesConverted []os.DirEntry
type Info struct {
	filename string
	err error
}
/*
Convert Mp4 to Apple Hls
//command := "ffmpeg -i fileInput -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls /Users/trinhdt2/learn/golang-techmaster/golang/lesson06/homeworks/data/testhls/index.m3u8"
*/
func ConvertMp4ToAppleHls(fileInput, frameSize string, partVideoTimeLength int64, fileOutput string, wg *sync.WaitGroup, ch chan<- Info) {
	command := fmt.Sprintf("%s%s%s%s%s%d%s%s", "ffmpeg -i ", fileInput, " -profile:v baseline -level 3.0 -s ", frameSize, " -start_number 0 -hls_time ", partVideoTimeLength, " -hls_time 10 -hls_list_size 0 -f hls ", fileOutput)
	err := utils.ExecCommand(command)
	ch <- Info{
		filename: fileInput,
		err: err,
	}
	wg.Done()
}


func Convert(inputPath string, outputPath string, partVideoTimeLength int64) {
	var err error
	var entriesNotYetConverted []os.DirEntry
	entriesConverted, entriesNotYetConverted, err = VideosAreNotYetConverted(inputPath, outputPath, entriesConverted)
	if err != nil {
		log.Printf(err.Error())
	}
	ch := make(chan Info)
	handle_convert_error(ch)


	if len(entriesNotYetConverted) == 0 {
		fmt.Println("Tất cả videos đã được chuyển đổi định dạng trước đây hoặc thư mục đầu ra đã tồn tại file")
	}else{
		wg := sync.WaitGroup{}
		for _, e := range entriesNotYetConverted {
			info, _ := e.Info()
			fmt.Println(info.ModTime())
			if strings.Compare(filepath.Ext(e.Name()), ".mp4") == 0 {
				wg.Add(1)
				go ConvertMp4ToAppleHls(inputPath+e.Name(), "1920X1080", partVideoTimeLength, outputPath+e.Name(), &wg, ch)
			}
		}
		wg.Wait()
		fmt.Println("Chuyển đổi định dạng file thành công")
	}
}

func Convert2(inputPath string, outputPath string, partVideoTimeLength int64, entriesNotYetConverted []os.DirEntry) {
	ch := make(chan Info)
	handle_convert_error(ch)
	if len(entriesNotYetConverted) == 0 {
		fmt.Println("Tất cả videos đã được chuyển đổi định dạng trước đây hoặc thư mục đầu ra đã tồn tại file")
	}else{
		wg := sync.WaitGroup{}
		for _, e := range entriesNotYetConverted {
			info, _ := e.Info()
			fmt.Println(info.ModTime())
			if strings.Compare(filepath.Ext(e.Name()), ".mp4") == 0 {
				wg.Add(1)
				go ConvertMp4ToAppleHls(inputPath+e.Name(), "1920X1080", partVideoTimeLength, outputPath+e.Name(), &wg, ch)
			}
		}
		wg.Wait()
		fmt.Println("Chuyển đổi định dạng file thành công")
	}
}

/*
Receive channel để xử lý thông báo lỗi
*/
func handle_convert_error(ch <-chan Info) {
	go func() {
		for {
			convert_result, more := <-ch //khi không còn dữ liệu, more sẽ false
			if convert_result.err != nil {
				fmt.Println(convert_result.err.Error())
			} else {
				fmt.Println("convert successfully " + convert_result.filename)
			}

			if !more {
				fmt.Println("No more result")
				return //thoát khỏi go routine
			}
		}
	}()
}

//Checking the videos is not convert yet
func VideosAreNotYetConverted(inputPath, outputPath string, entriesConverted []os.DirEntry) ([]os.DirEntry, []os.DirEntry, error) {
	var entriesNotYetConverted []os.DirEntry
	inputEntries, err := os.ReadDir(inputPath)
	if err != nil {
		return entriesConverted, entriesNotYetConverted, err
	}
	if entriesConverted == nil {
		outputEntries, err := os.ReadDir(outputPath)
		if err != nil {
			return entriesConverted, entriesNotYetConverted, err
		}
		for _, inputEntry := range inputEntries {
			if  len(outputEntries) == 0 {
				entriesNotYetConverted = append(entriesNotYetConverted, inputEntry)
			}else{
				if !utils.FileExists(outputPath+inputEntry.Name()) {
					entriesNotYetConverted = append(entriesNotYetConverted, inputEntry)
				}
			}
		}
	}else{
		for _, inputEntry := range inputEntries {
			if !utils.CheckEntryExsisted(entriesConverted, inputEntry) {
				entriesNotYetConverted = append(entriesNotYetConverted, inputEntry)
			}
		}
	}
	return entriesConverted, entriesNotYetConverted, nil
}
//Checking the videos is not convert yet
func VideosAreNotYetConverted2(inputPath, outputPath string, entriesConverted []os.DirEntry) ([]os.DirEntry, []os.DirEntry, error) {
	var entriesNotYetConverted []os.DirEntry
	inputEntries, err := os.ReadDir(inputPath)
	if err != nil {
		return entriesConverted, entriesNotYetConverted, err
	}
	if entriesConverted == nil {
		outputEntries, err := os.ReadDir(outputPath)
		if err != nil {
			return entriesConverted, entriesNotYetConverted, err
		}
		for _, inputEntry := range inputEntries {
			if  len(outputEntries) == 0 {
				entriesNotYetConverted = append(entriesNotYetConverted, inputEntry)
			}else{
				if !utils.FileExists(outputPath+inputEntry.Name()) {
					entriesNotYetConverted = append(entriesNotYetConverted, inputEntry)
				}
			}
		}
	}else{
		for _, inputEntry := range inputEntries {
			if !utils.CheckEntryExsisted(entriesConverted, inputEntry) {
				entriesNotYetConverted = append(entriesNotYetConverted, inputEntry)
			}
		}
	}
	return entriesConverted, entriesNotYetConverted, nil
}