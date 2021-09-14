package calendar

import "testing"
//go test -bench .
/*Benchmark_Calendar
*/
func Benchmark_Calendar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Calendar()
	}
}
/*Benchmark_PrintCalendar
 */
func Benchmark_PrintCalendar(b *testing.B) {
	cal := [5][7]int{}
	cal[0][0] = 0
	cal[0][1] = 0
	cal[0][2] = 0
	cal[0][3] = 1
	cal[0][4] = 2
	cal[0][5] = 3
	cal[0][6] = 4
	cal[1][0] = 5
	cal[1][1] = 6
	cal[1][2] = 7
	cal[1][3] = 8
	cal[1][4] = 9
	cal[1][5] = 10
	cal[1][6] = 11
	cal[2][0] = 12
	cal[2][1] = 13
	cal[2][2] = 14
	cal[2][3] = 15
	cal[2][4] = 16
	cal[2][5] = 17
	cal[2][6] = 18
	cal[3][0] = 19
	cal[3][1] = 20
	cal[3][2] = 21
	cal[3][3] = 22
	cal[3][4] = 23
	cal[3][5] = 24
	cal[3][6] = 25
	cal[4][0] = 26
	cal[4][1] = 27
	cal[4][2] = 28
	cal[4][3] = 29
	cal[4][4] = 30
	cal[4][5] = 0
	cal[4][6] = 0
	for i := 0; i < b.N; i++ {
		PrintCalendar(cal)
	}
}