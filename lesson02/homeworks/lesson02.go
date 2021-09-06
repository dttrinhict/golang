package main

import (
	"fmt"
	"github.com/dttrinhict/golang/lesson02/homeworks/slice"
)
func main() {
	//Bài 1 Viết function tìm ra số lớn thứ nhì trong mảng các số.
	args1 := []float64{1,9,2,8.4,3,7,4,6,5,8.8,4}
	fmt.Println("Số lớn thứ nhì trong mảng các số ", args1)
	max2 := slice.Max2Numbers(args1)
	fmt.Println(max2)

	//Bài 2 Cho 1 mảng các chuỗi. Viết function lọc ra các phần tử có độ dài lớn nhất.

	args := []string{"1", "22", "333", "22", "333", "444", "55", "555"}
	fmt.Println("Các phần tử có độ dài lớn nhất trong mảng ", args)
	result := slice.FindMaxLengthElement(args)
	for _, value := range result {
		fmt.Printf("%v ",value)
	}
	fmt.Println()
	//Bài 3 Viết function remove những phần tử bị trùng nhau trong mảng
	args3 := []float64{1, 22, 333, 22, 333, 444, 55, 444, 555}
	fmt.Println("Remove những phần tử bị trùng nhau trong mảng:", args3)
	rs3 := slice.RemoveDuplicates(args3)
	fmt.Println(rs3)

	/*Bài 4 Một nhân viên trong công ty bao gồm các thuộc tính sau : Tên, Hệ số lương, Tiền trợ cấp
	Tạo 1 mảng nhân viên (số lượng tuỳ ý) và thực hiện các chức năng sau:
	*/
	var nhanViens = []slice.NhanVien{
		{"CBA", 1.1, 1000},
		{"CAC", 1.2, 1000},
		{"DBC", 1.3, 100},
		{"BAC", 1.1, 1000000},
		{"ABC", 1.3, 100},
	}
	//Sắp xếp tên nhân viên tăng dần theo bảng chữ cái
	fmt.Println("Sắp xếp theo tên")
	nvs := slice.SortByName(nhanViens)
	for _, v := range nvs {
		fmt.Println(slice.ToString(&v))
	}
	//Sắp xếp nhân viên theo mức lương giảm dần (lương = Hệ số lương * 1.500.000 + Tiền trợ cấp)
	fmt.Println("Sắp xếp theo lương")
	nvs = slice.SortBySalary(nhanViens)
	for _, v := range nvs {
		fmt.Println(slice.ToString(&v))
	}
	//Lấy ra danh sách nhân viên có mức lương lớn thứ 2 trong mảng nhân viên
	fmt.Println("Danh sách nhân viên có mức lương cao thứ 2")
	nvs = slice.GetListBySalary(nhanViens)
	for _, v := range nvs {
		fmt.Println(slice.ToString(&v))
	}
}