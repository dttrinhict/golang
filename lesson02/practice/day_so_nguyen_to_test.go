package main

import "testing"

/*
Hướng dẫn kiểm tra tốc độ thực thi hàm
go test -bench .
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
