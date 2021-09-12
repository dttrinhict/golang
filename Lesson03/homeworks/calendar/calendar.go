package calendar

import (
	"fmt"
	"log"
	"time"
)

/*Lịch của tháng hiện tại
func Calendar() (calendar [5][7]int){
	now := time.Now() // Time now
	dayOfWeek := now.Weekday() // Ngày hiện tại trong tuần (thứ)
	dateOfMonth := now.Day() // Ngày hiện tại trong tháng
	month := now.Month() // Tháng hiện tại trong năm
	year := now.Year() // Năm hiện tại
	numDayOfMonth := NumDayOfMonth(month, year) // Số này của tháng hiện tại
	firstDayOfMonth := now.Weekday() - (time.Weekday(dateOfMonth%7)) + 1 // Ngày trong tuần của ngày 1 của tháng

	numDays :=  numDayOfMonth + int(firstDayOfMonth) -1 // Số ngày của tháng công ngày bù để tính lịch
	//calendar := [5][7]int{} // Khai báo calendar là một mảng số nguyên 2 chiều
	for j := 0; j <= numDays ; j++ {
		a := j%7 // ngày trong tuần
		b := j/7 // tuần trong tháng
		if j+1 - int(dayOfWeek - time.Weekday(dateOfMonth%7) + 1) <= 0 {
			calendar[b][a] = 0
		}else{
			calendar[b][a] = j+1 - int(dayOfWeek - time.Weekday(dateOfMonth%7) + 1)
		}
	}
	return calendar
}
*/
func Calendar() (calendar [5][7]int){
	now := time.Now() // Time now
	//dayOfWeek := now.Weekday() // Ngày hiện tại trong tuần (thứ)
	dayOfMonth := now.Day() // Ngày hiện tại trong tháng
	month := now.Month() // Tháng hiện tại trong năm
	year := now.Year() // Năm hiện tại
	numDayOfMonth := NumDayOfMonth(month, year) // Số này của tháng hiện tại
	if numDayOfMonth == 0 {
		log.Printf("Error: Có lỗi về giờ của hệ thống")
		return
	}
	firstDayOfMonth := 7 - (time.Weekday(dayOfMonth%7)) + 1 // Ngày trong tuần của ngày 1 của tháng
	numDays :=  numDayOfMonth + int(firstDayOfMonth) -1 // Số ngày của tháng công ngày bù để tính lịch
	for j := 0; j <= numDays ; j++ {
		a := j%7 // ngày trong tuần
		b := j/7 // tuần trong tháng
		if j < int(firstDayOfMonth) {
			calendar[b][a] = 0
		}else{
			calendar[b][a] = j+1 - int(firstDayOfMonth)
		}
	}
	return calendar
}

/* In ra lich của tháng hiện tại
https://golangbyexample.com/print-output-text-color-console/
func PrintCalendar(calendar [5][7]int)  {
	//colorYellow := "\033[33m"
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	for j:=0; j<=5; j++{
		for i := 0 ; i <= 6; i++ {
			if j == 0 {
				temp := time.Weekday(i).String()
				fmt.Printf("\t%v", temp[0:3])
			}else{
				if calendar[j-1][i] == 0 {
					fmt.Printf("\t")
				}else{
					if calendar[j-1][i] == time.Now().Day() {
						fmt.Printf("\t%v[%v]%v", string(colorGreen), calendar[j-1][i], string(colorReset))
					}else{
						fmt.Printf("\t%v", calendar[j-1][i])
					}
				}
			}
		}
		fmt.Println()
	}
}
*/
func PrintCalendar(calendar [5][7]int)  {
	colorYellow := "\033[33m"
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	now := time.Now()
	dayOfWeek := now.Weekday() // Ngày hiện tại trong tuần (thứ)
	dateOfMonth := now.Day() // Ngày hiện tại trong tháng
	month := now.Month() // Tháng hiện tại trong năm
	year := now.Year() // Năm hiện tại
	fmt.Println("\tToday is", string(colorYellow),", ", dayOfWeek, dateOfMonth, month, year, (string(colorReset)) )
	for j:=0; j<=5; j++{
		for i := 0 ; i <= 6; i++ {
			if j == 0 {
				temp := time.Weekday(i).String()
				fmt.Printf("\t%v", temp[0:3])
			}else{
				if calendar[j-1][i] == 0 {
					fmt.Printf("\t")
				}else{
					if calendar[j-1][i] == time.Now().Day() {
						fmt.Printf("\t%v[%v]%v", string(colorGreen), calendar[j-1][i], string(colorReset))
					}else{
						fmt.Printf("\t%v", calendar[j-1][i])
					}
				}
			}
		}
		fmt.Println()
	}
}

func NumDayOfMonth(month time.Month, year int) (numday int)  {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		numday = 31
	case 4,6,9,11:
		numday = 30
	case 2:
		if year % 4 == 0 {
			numday = 29
		}else {
			numday = 28
		}
	default:
		numday = 0
	}
	return numday
}

func DayOfWeek() (day time.Weekday) {
	return time.Now().Weekday() - 5
}

func GetDateOfMonnth() (dateOfMonth int) {
	return time.Now().Day() - 1
}