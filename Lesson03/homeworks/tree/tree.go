package tree

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type TreeDir struct {
	value interface{}
	Node *TreeDir
}

func NewTree() TreeDir {
	return TreeDir{}
}

func ListElementOfDir(root string) ([]string,string, error) {
	if !strings.HasSuffix(root, "/") {
		root = root + "/"
	}
	var files = []string{}
	dirElements, err := ioutil.ReadDir(root)
	for _, e := range dirElements {
		if e.IsDir() {
			curentDir := root + e.Name()
			defer func(root string) {
				if fs, _,e := ListElementOfDir(root); e != nil {
					log.Println(e.Error())
				}else{
					for _, f := range fs {
						files = append(files, f)
					}
				}
			}(curentDir)
		}else{
			file := root + e.Name()
			files = append(files, file)
			fmt.Println(file)
		}
	}
	return files,root, err
}

/* In ra cay thu muc
*/
func ListElementOfFolder(root string) ([]string,string, error) {
	if !strings.HasSuffix(root, "/") {
		root = root + "/"
	}
	var oldDir string
	var files = []string{}
	dirElements, err := ioutil.ReadDir(root)
	for _, e := range dirElements {
		if e.IsDir() {
			curentDir := root + e.Name()
			oldDir = root
			defer func(root string, oldDir string) {
				if fs, _,e := ListElementOfFolder(root); e != nil {
					log.Println(e.Error())
				}else{
					for _, f := range fs {
						files = append(files, f)
					}
				}
			}(curentDir, oldDir)

		}else{
			file := root + e.Name()
			files = append(files, file)
			rootDirSplit := strings.Split(root,"/")
			rootDeep := len(rootDirSplit)
			i := 0
			if strings.Compare(oldDir, root) == 0 && oldDir != ""{
				fmt.Printf("f")
				for i < rootDeep {
					fmt.Printf(" ")
					i++
				}
				fmt.Println(e.Name())
			}else if oldDir == "" {
				fmt.Printf("d")
				for i < rootDeep -1 {
					fmt.Printf("-")
					i++
				}
				fmt.Println(rootDirSplit[rootDeep - 2])
				j := 0
				fmt.Printf("f")
				for j < rootDeep {
					fmt.Printf(" ")
					j++
				}
				fmt.Println(e.Name())
			}
		}
	}
	return files,root, err
}