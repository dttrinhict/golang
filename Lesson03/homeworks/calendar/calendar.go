package calendar

import (
	"fmt"
	"log"
	"time"
)

type Cal struct {
	CurrentTime time.Time
	BodyCal map[int][]int
}

func InitCalendar() *Cal {
	return &Cal{
		CurrentTime: time.Now(),
		BodyCal: map[int][]int{},
	}
}

const colorYellow = "\033[33m"
const colorReset = "\033[0m"
const colorGreen = "\033[32m"

/**Code refactor
/* Get weekday of first day of month
 */
func (c Cal) FirstDayOfMonth2() (firtDayOfMonth time.Weekday) {
	month := c.CurrentTime.Month()
	year := c.CurrentTime.Year()
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	return date.Weekday()
}

/* Calendar of current month
 */
func (c *Cal) Calendar2() Cal {
	calendar := map[int][]int{}
	now := time.Now()
	firstDayOfMonth := c.FirstDayOfMonth2()
	numDaysOfMonth := NumDayOfMonth(now.Month(), now.Year())
	j := 1
	for j <= int(firstDayOfMonth) {
		calendar[0] = append(calendar[0], 0)
		j++
	}
	i := 1
	for i <= numDaysOfMonth {
		week := (i + int(firstDayOfMonth) - 1) / 7
		calendar[week] = append(calendar[week], i)
		i++
	}
	return Cal{
		CurrentTime: now,
		BodyCal: calendar,
	}
}

/* Show current month on terminal
*/
func (c *Cal) PrintCalendar2(cal Cal) {
	calBody := cal.BodyCal
	fmt.Println("\tToday is", string(colorYellow), ", ", c.CurrentTime.Weekday(), c.CurrentTime.Day(), c.CurrentTime.Month(), c.CurrentTime.Year(), (string(colorReset)))
	for i := 0; i <= int(time.Saturday); i++ {
		temp := time.Weekday(i).String()
		fmt.Printf("\t%v%v%v", string(colorYellow), temp[:3], string(colorReset))
	}
	fmt.Println()
	for j:=0 ; j<=len(calBody); j++{
		for _, c := range calBody[j] {
			PrintDay(c)
		}
		fmt.Println()
	}
}

/* Format print a date
*/
func PrintDay(day int) {
	today := time.Now().Day()
	if day == 0 {
		fmt.Printf("\t%s", "")
	}else{
		if day == today {
			fmt.Printf("\t%v[%v]%v", string(colorGreen), day, string(colorReset))
		}else{
			fmt.Printf("\t%v", day)
		}
	}
}

/* Xác định số ngày trong tháng
 */
func NumDayOfMonth(month time.Month, year int) (numday int) {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		numday = 31
	case 4, 6, 9, 11:
		numday = 30
	case 2:
		if year%4 == 0 {
			numday = 29
		} else {
			numday = 28
		}
	default:
		numday = 0
	}
	return numday
}






/* Bí code
*/
/*Lịch của tháng hiện tại
 */
func Calendar() (calendar [6][7]int) {
	now := time.Now() // Time now
	//now := time.Date(2021,13, 2, 0, 0, 0, 0, time.Local)
	dayOfWeek := now.Weekday() // Ngày hiện tại trong tuần (thứ)
	dayOfMonth := now.Day() // Ngày hiện tại trong tháng
	month := now.Month()                        // Tháng hiện tại trong năm
	year := now.Year()                          // Năm hiện tại
	numDayOfMonth := NumDayOfMonth(month, year) // Số này của tháng hiện tại
	if numDayOfMonth == 0 {
		log.Printf("Error: Có lỗi về giờ của hệ thống")
		return
	}
	firstDayOfMonth := FirstDayOfMonth(dayOfWeek, dayOfMonth)
	//firstDayOfMonth := int(FirstDayOfMonth2(now))
	numDays := numDayOfMonth + firstDayOfMonth - 1 // Số ngày của tháng công ngày bù để tính lịch
	for j := 0; j <= numDays; j++ {
		a := j % 7 // ngày trong tuần
		b := j / 7 // tuần trong tháng
		if j < firstDayOfMonth {
			calendar[b][a] = 0
		} else {
			calendar[b][a] = j + 1 - firstDayOfMonth
		}
	}
	return calendar
}

/* In ra lich của tháng hiện tại
 */
func PrintCalendar(calendar [6][7]int) {
	now := time.Now()
	dayOfWeek := now.Weekday() // Ngày hiện tại trong tuần (thứ)
	dateOfMonth := now.Day()   // Ngày hiện tại trong tháng
	month := now.Month()       // Tháng hiện tại trong năm
	year := now.Year()         // Năm hiện tại
	fmt.Println("\tToday is", string(colorYellow), ", ", dayOfWeek, dateOfMonth, month, year, (string(colorReset)))
	for j := 0; j <= 5; j++ {
		for i := 0; i <= 6; i++ {
			if j == 0 {
				temp := time.Weekday(i).String()
				fmt.Printf("\t%v", temp[0:3])
			} else {
				if calendar[j-1][i] == 0 {
					fmt.Printf("\t")
				} else {
					if calendar[j-1][i] == time.Now().Day() {
						fmt.Printf("\t%v[%v]%v", string(colorGreen), calendar[j-1][i], string(colorReset))
					} else {
						fmt.Printf("\t%v", calendar[j-1][i])
					}
				}
			}
		}
		fmt.Println()
	}
}

func FirstDayOfMonth(dayOfWeek time.Weekday, dayOfMonth int) (firstDayOfMonth int) {
	firstDayOfMonth = int(dayOfWeek) - (dayOfMonth % 7) + 1
	if firstDayOfMonth < 0 {
		firstDayOfMonth = firstDayOfMonth + 7
	} else if firstDayOfMonth < 7 {
		firstDayOfMonth = firstDayOfMonth
	} else {
		firstDayOfMonth = 0
	}
	return firstDayOfMonth
}