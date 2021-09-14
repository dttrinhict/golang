package main

import (
	"github.com/dttrinhict/golang/lesson03/homeworks/calendar"
)

func main()  {
	cal2 := calendar.InitCalendar()
	c := cal2.Calendar2()
	cal2.PrintCalendar2(c)
}
