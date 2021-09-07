package slice

import (
	"strings"
	"testing"
)

func Test_SortByName(t *testing.T) {
	var tests = []struct {
		name string
		agrs []NhanVien
		want []NhanVien
	}{
		{
			"Sorted-by-name",
			[]NhanVien{
			{"CBA", 1.1, 1000},
			{"CAC", 1.2, 1000},
			{"DBC", 1.3, 100},
			{"BAC", 1.1, 1000000},
			{"ABC", 1.3, 100},
			},
		[]NhanVien{
			{"ABC", 1.3, 100},
			{"BAC", 1.1, 1000000},
			{"CAC", 1.2, 1000},
			{"CBA", 1.1, 1000},
			{"DBC", 1.3, 100},
			},
		},
	}
	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			ans := SortByName(tt.agrs)
			for i, v := range ans {
				if strings.Compare(v.Ten,tt.want[i].Ten) != 0 {
					t.Errorf("got %v, want %v", ans, tt.want)
				}
			}
		})
	}
}

func BenchmarkSortByName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SortByName([]NhanVien{
			{"CBA", 1.1, 1000},
			{"CAC", 1.2, 1000},
			{"DBC", 1.3, 100},
			{"BAC", 1.1, 1000000},
			{"ABC", 1.3, 100},
		})
	}
}

func Test_SortBySalary(t *testing.T) {
	var tests = []struct {
		name string
		agrs []NhanVien
		want []NhanVien
	}{
		{
			"Sorted-by-salary",
			[]NhanVien{
				{"CBA", 1.1, 1000},
				{"CAC", 1.2, 1000},
				{"DBC", 1.3, 100},
				{"BAC", 1.1, 1000000},
				{"ABC", 1.3, 100},
			},
			[]NhanVien{
				{"CBA", 1.1, 1000},
				{"CAC", 1.2, 1000},
				{"DBC", 1.3, 100},
				{"ABC", 1.3, 100},
				{"BAC", 1.1, 1000000},
			},
		},
	}
	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			ans := SortBySalary(tt.agrs)
			for i, v := range ans {
				if strings.Compare(v.Ten,tt.want[i].Ten) != 0 {
					t.Errorf("got %v, want %v", ans, tt.want)
				}
			}
		})
	}
}
func BenchmarkSortBySalary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SortBySalary([]NhanVien{
			{"CBA", 1.1, 1000},
			{"CAC", 1.2, 1000},
			{"DBC", 1.3, 100},
			{"BAC", 1.1, 1000000},
			{"ABC", 1.3, 100},
		})
	}
}

func Test_GetListBySalary(t *testing.T) {
	var tests = []struct {
		name string
		agrs []NhanVien
		want []NhanVien
	}{
		{
			"get-by-salary",
			[]NhanVien{
				{"CBA", 1.1, 1000},
				{"CAC", 1.2, 1000},
				{"DBC", 1.3, 100},
				{"BAC", 1.1, 1000000},
				{"ABC", 1.3, 100},
			},
			[]NhanVien{
				{"ABC", 1.3, 100},
				{"DBC", 1.3, 100},
			},
		},
		{
			"get-by-salary-return-nil",
			[]NhanVien{
				{"DBC", 1.3, 100},
				{"ABC", 1.3, 100},
			},
			[]NhanVien{},
		},
	}
	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			ans := GetListBySalary(tt.agrs)
			for i, v := range ans {
				if strings.Compare(v.Ten,tt.want[i].Ten) != 0 {
					t.Errorf("got %v, want %v", ans, tt.want)
				}
			}
		})
	}
}
func BenchmarkGetListBySalary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetListBySalary([]NhanVien{
			{"CBA", 1.1, 1000},
			{"CAC", 1.2, 1000},
			{"DBC", 1.3, 100},
			{"BAC", 1.1, 1000000},
			{"ABC", 1.3, 100},
		})
	}
}

func Test_Salary(t *testing.T) {
	var tests = []struct {
		name string
		agrs NhanVien
		want float64
	}{
		{
			"get-by-salary",
			NhanVien{
				"CBA", 1.1, 1000,
			},
			float64(1.1)*1500000+float64(1000),
		},
	}
	for _, tt := range tests {
		testname := tt.name
		t.Run(testname, func(t *testing.T) {
			got := Salary(tt.agrs)
			if got != tt.want {
				t.Errorf("got %v, want %v \n", got, tt.want)
			}
		})
	}
}
func BenchmarkSalary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Salary(NhanVien{
			"CBA",
			1.1,
			1000})
	}
}
