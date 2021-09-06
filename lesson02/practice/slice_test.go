package main

import "testing"

func Benchmark_removeSliceItemNotKeepOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeSliceItemNotKeepOrder([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}, 5)
	}
}
