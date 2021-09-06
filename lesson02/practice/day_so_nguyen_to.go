package main

import "math"

/*Xác định dãy số nguyên tố trong một dãy số nguyên dương N phần tử.
 */
func DaySoNguyenTo(n int64) (daysonguyento []int64) {
	var i int64
	for i = 2; i <= n; i++ {
		laSoNguyenTo := LaSoNguyenTo(i)
		if laSoNguyenTo {
			daysonguyento = append(daysonguyento, i)
		}
	}
	return daysonguyento
}

/* Kiểm tra số a có phải là số nguyên tố không
 */
func LaSoNguyenTo(a int64) bool {
	var j int64
	sqrtn := int64(math.Sqrt(float64(a)))
	for j = 2; j <= sqrtn; j++ {
		if a%j == 0 {
			return false
		}
	}
	return true
}
