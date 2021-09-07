package slice

import (
	"testing"
)
/*
Hướng dẫn kiểm tra tốc độ thực thi hàm
go test -bench .
*/
/* lesson02
*/

/* Unit test for Max2Numbers
*/
func Test_Max2Numbers(t *testing.T) {
	var tests = []struct {
		name string
		agrs []float64
		want float64
	}{
		{
			"max-2-number",
			[]float64{1,9,2,8.4,3,8.9,7,4,6,5,8.8,4},
			float64(8.9),
		},
	}
	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			got := Max2Numbers(tt.agrs)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

/*Benchmark for Max2Numbers function
 */
func Benchmark_Max2Numbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Max2Numbers([]float64{1,9,2,8.4,3,7,4,6,5,8.8,4})
	}
}

/* Unit test for FindMaxLengthElement
 */
func Test_FindMaxLengthElement(t *testing.T) {
	var tests = []struct {
		name string
		agrs []string
		want []string
	}{
		{
			"max-2-number",
			[]string{"1","22","333", "44", "555","4","444"},
			[]string{"333","555","444"},
		},
	}
	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			got := FindMaxLengthElement(tt.agrs)
			for i,v := range got {
				if v != tt.want[i] {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			}
		})
	}
}

/*Benchmark for FindMaxLengthElement function
 */
func Benchmark_FindMaxLengthElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FindMaxLengthElement([]string{"1","22","333", "44", "555","4","444"})
	}
}

/* Unit test for MaxLengthElements
 */
func Test_MaxLengthElements(t *testing.T) {
	var tests = []struct {
		name string
		maxLen int
		agrs []string
		want []string
	}{
		{
			"max-2-number",
			3,
			[]string{"1","22","333", "44", "555","4","444"},
			[]string{"333","555","444"},
		},
	}
	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			got := MaxLengthElements(tt.maxLen,tt.agrs)
			for i,v := range got {
				if v != tt.want[i] {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			}
		})
	}
}

/*Benchmark for MaxLengthElements function
 */
func Benchmark_MaxLengthElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MaxLengthElements(3,[]string{"1","22","333","555","444"})
	}
}

/* Unit test for RemoveItem
 */
func Test_RemoveItem(t *testing.T) {
	var tests = []struct {
		name string
		agrs []float64
		index int
		want []float64
	}{
		{
			"max-2-number",
			[]float64{1, 22, 333, 44, 555, 4, 444},
			3,
			[]float64{1, 22, 333, 555, 4, 444},
		},
	}
	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			got := RemoveItem(tt.agrs, tt.index)
			for i,v := range got {
				if v != tt.want[i] {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			}
		})
	}
}

/*Benchmark for RemoveItem function
 */
func Benchmark_RemoveItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RemoveItem([]float64{1, 22, 333, 22, 333, 444, 55, 444, 555}, 4)
	}
}

/*Unit test for RemoveDuplicates
*/
func Test_RemoveDuplicates(t *testing.T) {
	var tests = []struct {
		name string
		agrs []float64
		want []float64
	}{
		{
			"max-2-number",
			[]float64{1, 22, 333, 22, 333, 444, 55, 444, 555},
			[]float64{1, 22, 333, 444, 55, 555},
		},
	}
	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			got := RemoveDuplicates(tt.agrs)
			for i,v := range got {
				if v != tt.want[i] {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			}
		})
	}
}

/*Benchmark for RemoveDuplicates function
*/
func Benchmark_RemoveDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RemoveDuplicates([]float64{1, 22, 333, 22, 333, 444, 55, 444, 555})
	}
}
/*Lesson01
 */


/*Unit test for DaySoNguyenTo
 */
func Test_DaySoNguyenTo(t *testing.T) {
	var tests = []struct {
		name string
		n int64
		want []int64
	}{
		{
			"Day_so_nguyen_to",
			20,
			[]int64{2, 3, 5, 7, 11, 13, 17, 19},
		},
	}
	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			got := DaySoNguyenTo(tt.n)
			for i,v := range got {
				if v != tt.want[i] {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			}
		})
	}
}
/* Benchmark for DaySoNguyenTo
 */
func Benchmark_DaySoNguyenTo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = DaySoNguyenTo(1000)
	}
}

/*Unit test for DaySoNguyenTo
 */
func Test_LaSoNguyenTo(t *testing.T) {
	var tests = []struct {
		name string
		so int64
		want bool
	}{
		{
			"Khong_la_so_nguyen_to",
			20,
			false,
		},
		{
			"La_so_nguyen_to",
			19,
			true,
		},
	}
	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			got := LaSoNguyenTo(tt.so)
			if !(got && tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
/* Benchmark for LaSoNguyenTo
 */
func Benchmark_LaSoNguyenTo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = LaSoNguyenTo(7)
	}
}
