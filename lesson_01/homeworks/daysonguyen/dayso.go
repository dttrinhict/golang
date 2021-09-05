package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Hãy nhập vào số lớn nhất của dãy số: ")
	str, _ := reader.ReadString('\n')
	so, _ := ChuyenChuoiSangInt(str)
	for so >= 100 {
		fmt.Printf("Hãy nhập vào số lớn nhất của dãy số < 100: ")
		str, _ := reader.ReadString('\n')
		so, _ = ChuyenChuoiSangInt(str)
	}
	daySoNguyenTo := DaySoNguyenTo(so)
	fmt.Printf("Các số nguyên tố trong dãy số từ 1 đến %v là\n",so )
	for _, value := range daySoNguyenTo {
		fmt.Print(value, " ")
	}
}

/* Tìm số nguyên tố trong một dãy số tự nhiên
func DaySoNguyenTo(so int64) []int64 {
	var daySoNguyenTo []int64
	var i int64
	for i = 1; i <= so; i++ {
		lasonguyento := true
		var j int64 = 2
		for j < i {
			if i%j == 0 {
				lasonguyento = false
				break
			}
			j++
		}
		if lasonguyento {
			daySoNguyenTo = append(daySoNguyenTo, i)
		}
	}
	return daySoNguyenTo
}
*/
func DaySoNguyenTo(so int64) []int64 {
	var daySoNguyenTo []int64
	var i int64
	for i = 1; i <= so; i++ {
		lasonguyento := true
		var j int64 = 2
		for j < i {
			if i%j == 0 {
				lasonguyento = false
				break
			}
			j++
		}
		if lasonguyento {
			daySoNguyenTo = append(daySoNguyenTo, i)
		}
	}
	return daySoNguyenTo
}

/* Chuyển dữ liệu chuỗi đầu vào sang int64
 */
func ChuyenChuoiSangInt(str string) (int64, error) {
	return strconv.ParseInt(strings.Trim(str, "\n"), 10, 8)
}
