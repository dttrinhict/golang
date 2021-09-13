package main

import (
	"bufio"
	"fmt"
	"github.com/dttrinhict/golang/lesson03/homeworks/tree"
	"log"
	"os"
	"strings"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter Folder Path: ")
	k, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err.Error())
	}
	dir := strings.Trim(k, "\n")
	//dir := "./lesson03/homeworks"
	fmt.Println(dir)
	_, _, errList := tree.ListElementOfFolder(dir)
	if errList != nil {
		log.Println(err.Error())
	}
}
