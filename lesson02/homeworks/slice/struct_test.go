package slice

import "testing"

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

func BenchmarkSalary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Salary(NhanVien{
			"CBA",
			1.1,
			1000})
	}
}
