package slice

import (
	"fmt"
	"sort"
)

type NhanVien struct {
	Ten string
	HeSo float64
	TroCap float64
}

func ToString(nv *NhanVien) string {
	return fmt.Sprintf("Tên: %v, Hệ Số: %v, Trợ Cấp: %v", nv.Ten, nv.HeSo, nv.TroCap)
}

/*Sắp xếp tên nhân viên tăng dần theo bảng chữ cái
https://stackoverflow.com/questions/36122668/how-to-sort-struct-with-multiple-sort-parameters
*/
func SortByName(nhanViens []NhanVien) ([]NhanVien) {
	sort.Slice(nhanViens, func(i, j int) bool {
		return nhanViens[i].Ten < nhanViens[j].Ten
	})
	return nhanViens
}

/*
Sắp xếp nhân viên theo mức lương giảm dần (lương = Hệ số lương * 1.500.000 + Tiền trợ cấp)
https://yourbasic.org/golang/how-to-sort-in-go/
 */
func SortBySalary(nhanViens []NhanVien) ([]NhanVien) {
	sort.Slice(nhanViens, func(i, j int) bool {
		return Salary(nhanViens[i]) < Salary(nhanViens[j])
	})
	return nhanViens
}

/*Lấy ra danh sách nhân viên có mức lương lớn thứ 2 trong mảng nhân viên
 */
func GetListBySalary(nhanViens []NhanVien) (result []NhanVien) {
	temp := SortBySalary(nhanViens)
	j := 0
	n := len(temp)-1
	for i := n-1; i >=0; i-- {
		if Salary(temp[i]) < Salary(temp[n]) {
			result = append(result, temp[i])
			j = i
		}
		for k := j - 1; k>=0; k-- {
			if Salary(temp[k]) == Salary(temp[j]) {
				result = append(result, temp[k])
			}else{
				return result
			}
		}
	}
	return result
}

/* tính lương
 */
func Salary(nv NhanVien) float64 {
	return nv.HeSo*1500000+nv.TroCap
}