package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main()  {
	var str string
	fmt.Println("Gỉai phương trình bậc 2 dạng ax2 + bx + c = 0")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nhập vào hệ số a: ")
	str, _ = reader.ReadString('\n')
	a, _ := ChuyenChuoiSangFloat64(str)
	fmt.Print("Nhập vào hệ số b: ")
	str, _ = reader.ReadString('\n')
	b, _ := ChuyenChuoiSangFloat64(str)
	fmt.Print("Nhập vào hệ số c: ")
	str, _ = reader.ReadString('\n')
	c, _ := ChuyenChuoiSangFloat64(str)
	x1, x2, err := GiaiPhuongTrinhBac2(a, b, c)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}else{
		fmt.Printf("Các nghiệm của phương trình bậc 2 %vx2 + %vx + %v = 0 là: x1 = %v, x2 = %v\n", a, b, c, x1, x2)
	}
}

/* giải phương trình bậc 2 ax2 + bx + c = 0
nếu a, b = 0 không thành lập phương trình
nếu a = 0, b != 0 thành lập phương trình bậc nhất có một nghiệm là -c/b
nếu a,b,c != 0 thành lập phương trình bậc 2, giải phương trình bậc 2
- tính delta = b*b - 4*a*c
- nếu delta < 0 phương trình vô nghiệm
- nếu delta == 0 phương trình có nghiệm kép là -b/(2*a)
- nếu delta > 0 phương trình có 2 nghiệm
x1 = (-b + sqrt(delta))/(2*a)
x2 = (-b - sqrt(delta))/(2*a)
*/
func GiaiPhuongTrinhBac2(a,b,c float64) (float64, float64, error)  {
	if a == 0 {
		if b == 0 {
			return 0,0,errors.New("không thành lập phương trình")
		}
		if b != 0 {
			return -c/b, -c/b, nil
		}
	}
	delta := b*b - 4*a*c
	if delta < 0 {
		return 0,0, errors.New("Phương trình vô nghiệm")
	}
	if delta == 0 {
		return -b/(2*a), -b/(2*a), nil
	}
	return (-b+math.Sqrt(delta))/(2*a),(-b-math.Sqrt(delta))/(2*a),nil
}

/* Chuyển dữ liệu chuỗi đầu vào sang float64
*/
func ChuyenChuoiSangFloat64(str string) (float64, error) {
	return strconv.ParseFloat(strings.Trim(str,"\n"), 64)
}