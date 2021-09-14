package calendar

import (
	"testing"
	"time"
)

//go test -bench .

var cal = InitCalendar()
/* Benchmark for FirstDayOfMonth2 method of Cal struct
*/
func BenchmarkCal_FirstDayOfMonth2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cal.FirstDayOfMonth2()
	}
}

/*Benchmark for NumDayOfMonth func
*/
func BenchmarkNumDayOfMonth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NumDayOfMonth(time.September, 2021)
	}
}

/*Benchmark for Calendar2 method of Cal struct
 */
func BenchmarkCal_Calendar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cal.Calendar2()
	}
}

/*Benchmark for PrintCalendar2 method
*/
//func BenchmarkCal_PrintCalendar2(b *testing.B) {
//	c := cal.Calendar2()
//	for i := 0; i < b.N; i++ {
//		c.PrintCalendar2(c)
//	}
//}

/*Benchmark for PrintDay function
*/
//func BenchmarkPrintDay(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		PrintDay(4)
//	}
//}


/*Benchmark_Calendar
*/
func Benchmark_Calendar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Calendar()
	}
}