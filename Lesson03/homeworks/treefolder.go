package main

import (
	//"bufio"
	"fmt"
	"github.com/dttrinhict/golang/lesson03/homeworks/treefolder"
	"os"

	//"log"
	//"os"
	//"strings"
)

func main()  {
	//os.Args trả về mảng string gồm path của file chạy và các arguments
	if len(os.Args) < 2 {
		fmt.Println("usage: ./tree <path>")
		os.Exit(1)
	}
	fmt.Println(os.Args)
	dir := os.Args[1]
	treefolder.PrintDirectory(dir, 0)
}
