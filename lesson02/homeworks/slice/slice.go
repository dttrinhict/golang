package slice

import (
	"math"
)

/* lesson02
*/

/* Tìm số lớn thứ 2 trong dãy số
Bài 1 Viết function tìm ra số lớn thứ nhì trong mảng các số.
Ví dụ: max2Numbers([2, 1, 3, 4]) => 3
 */
func Max2Numbers(numbers []float64) (max2 float64) {
	var max1 float64 = numbers[0]
	max2 = max1
	for _, value := range numbers {
		if max1 < value {
			max2 = max1
			max1 = value
		} else if max2 < value {
			max2 = value
		}
	}
	return max2
}

/*Bài 2 Cho 1 mảng các chuỗi. Viết function lọc ra các phần tử có độ dài lớn nhất.
Ví dụ: findMaxLengthElement["aba", "aa", "ad", "c", "vcd"] => ["aba", "vcd"]
*/
func FindMaxLengthElement(args []string) (result []string) {
	var maxLen int = 0
	var temp []string
	for _, value := range args {
		if maxLen <= len(value) {
			maxLen = len(value)
			temp = append(temp, value)
		}
	}
	return MaxLengthElements(maxLen, temp)
}

/* mảng của các phần tử có đội dài lới nhất
*/
func MaxLengthElements(maxLen int, args []string) (result []string)  {
	for _, value := range args {
		if len(value) == maxLen {
			result = append(result, value)
		}
	}
	return result
}

/*Bài 3 Viết function remove những phần tử bị trùng nhau trong mảng
Ví dụ: removeDuplicates([1,2,5,2,6,2,5]) => [1,2,5,6]
*/
func RemoveDuplicates(args []float64) (result []float64)  {
	result = args
	var index []int // chưa index của những phân tử bị trùng lặp
	for i, value := range result {
		for j := i + 1; j < len(result); j++ {
			if value == result[j] {
				index = append(index,j)
			}
		}
	}
	for k := len(index) - 1; k >= 0; k-- {
		result = RemoveItem(result, index[k])
	}
	return result
}

/*remove item has index
*/
func RemoveItem(args []float64, index int) (result []float64)  {
	result = append(result, args[:index]...)
	result = append(result, args[index+1:]...)
	return  result
}



/*Lesson01
*/

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