package slice

import "testing"
/*
Hướng dẫn kiểm tra tốc độ thực thi hàm
go test -bench .
*/
/* lesson02
*/
func Benchmark_Max2Numbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Max2Numbers([]float64{1,9,2,8.4,3,7,4,6,5,8.8,4})
	}
}
func Benchmark_FindMaxLengthElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FindMaxLengthElement([]string{"1","22","333", "44", "555","4","444"})
	}
}
func Benchmark_MaxLengthElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MaxLengthElements(3,[]string{"1","22","333","555","444"})
	}
}
func Benchmark_RemoveItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RemoveItem([]float64{1, 22, 333, 22, 333, 444, 55, 444, 555}, 4)
	}
}
func Benchmark_RemoveDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RemoveDuplicates([]float64{1, 22, 333, 22, 333, 444, 55, 444, 555})
	}
}
/*Lesson01
 */

func Benchmark_DaySoNguyenTo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = DaySoNguyenTo(1000)
	}
}
func Benchmark_LaSoNguyenTo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = LaSoNguyenTo(7)
	}
}
