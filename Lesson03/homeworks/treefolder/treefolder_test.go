package treefolder

import "testing"

func Benchmark_PrintDirectory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintDirectory("./", 0)
	}
}

func Benchmark_PrintListing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintListing("./a/b/c/d/s/e", 6)
	}
}