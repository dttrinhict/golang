package treefolder

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type TreeDir struct {
	value interface{}
	Node  *TreeDir
}

func NewTree() TreeDir {
	return TreeDir{}
}

func ListElementOfDir(root string) ([]string, string, error) {
	if !strings.HasSuffix(root, "/") {
		root = root + "/"
	}
	var files = []string{}
	dirElements, err := ioutil.ReadDir(root)
	for _, e := range dirElements {
		if e.IsDir() {
			curentDir := root + e.Name()
			defer func(root string) {
				if fs, _, e := ListElementOfDir(root); e != nil {
					log.Println(e.Error())
				} else {
					for _, f := range fs {
						files = append(files, f)
					}
				}
			}(curentDir)
		} else {
			file := root + e.Name()
			files = append(files, file)
			fmt.Println(file)
		}
	}
	return files, root, err
}

/* In ra cay thu muc
 */
func PrintListing(entry string, depth int) {
	indent := strings.Repeat("|   ", depth)
	fmt.Printf("%s|-- %s\n", indent, entry)
}

/* Bí quá đành phải copy trên mạng về chạy
thuật toán sử dụng đệ quy để duyệt
có kiểm tra symlink và show ra symlink
bản quyền thuộc về malicote https://github.com/malicote/go-tree
*
*/
func PrintDirectory(path string, depth int) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("error reading %s: %s\n", path, err.Error())
		return
	}
	foldername := strings.Split(strings.Trim(path,"/"),"/")//thêm phần này để get folder name
	PrintListing(foldername[len(foldername)-1], depth)
	for _, entry := range entries {
		if (entry.Mode() & os.ModeSymlink) == os.ModeSymlink {
			full_path, err := os.Readlink(filepath.Join(path, entry.Name()))
			if err != nil {
				fmt.Printf("error reading link: %s\n", err.Error())
			} else {
				PrintListing(entry.Name()+" -> "+full_path, depth+1)
			}
		} else if entry.IsDir() {
			PrintDirectory(filepath.Join(path, entry.Name()), depth+1)
		} else {
			PrintListing(entry.Name(), depth+1)
		}
	}
}
