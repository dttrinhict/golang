package main

import (
	"fmt"
	"github.com/dttrinhict/golang/lesson06/homeworks/converts"
	"log"
	"os"
	"time"
)
var inputPath = "/Users/trinhdt2/learn/golang-techmaster/golang/lesson06/homeworks/data/inputmp4/"
var outputPath = "/Users/trinhdt2/learn/golang-techmaster/golang/lesson06/homeworks/data/outhls/"

func main()  {
	//converts.Convert(inputPath, outputPath, 15)

	var entriesConverted []os.DirEntry
	var entriesNotYetConverted []os.DirEntry
	var err error
	for {
		fmt.Println(entriesConverted)
		entriesConverted, entriesNotYetConverted, err = converts.VideosAreNotYetConverted(inputPath, outputPath, entriesConverted)
		if err != nil {
			log.Printf(err.Error())
		}
		if len(entriesNotYetConverted) != 0 {
			converts.Convert2(inputPath, outputPath, 15, entriesNotYetConverted)
		}
		time.Sleep(30*time.Second)
	}
}
