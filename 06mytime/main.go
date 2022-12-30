package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()

	//The input to the Format method should always be "01-02-2006 15:04:03 Monday" for "month-day-year hour:minute:second weekday"
	fmt.Println(presentTime.Format("02-01-2006 15:04:03 Monday"))

	//time.Date() is used to create a date by passing data from the users
	//To create a date we have to pass data such as time.Date(year(as int), month(as time.monthName), day(as int), hour(as int), minute(as int), seconds(as int), nanoSeconds(as int))
	createdDate := time.Date(2020, time.August, 4, 10, 30, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006 15:04:03 Monday"))
}
