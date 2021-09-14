package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("Time now: %v\n", time.Now().Unix())
	time.Sleep(2 * time.Second)
	fmt.Printf("New time: %v\n", time.Now().Unix())

	t := time.Date(2020, time.August, 10, 20, 0, 0, 0, time.UTC)
	fmt.Printf("Time: %s\n", t)

	fmt.Print("********************************\n")
	now := time.Now()
	fmt.Printf("Now: %v\n", now)

	fmt.Printf("Day: %v\n", now.Day())
	fmt.Printf("Month: %v\n", now.Month())
	fmt.Printf("Year: %v\n", now.Year())

	fmt.Print("********************************\n")

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v,%v,%v,%v\n", tomorrow.Weekday(), tomorrow.Day(), tomorrow.Month(), tomorrow.Year())

	fmt.Print("********************************\n")

	/* Time compering*/
	xtime := time.Date(1999, 02, 22, 22, 0, 0, 0, time.UTC)

	fmt.Println(xtime.Before(now))
	fmt.Println(xtime.After(now))
	fmt.Println(xtime.Equal(now))

	diff := now.Sub(xtime)
	fmt.Println(diff)
	fmt.Println(diff.Hours())
	fmt.Print(diff.Minutes())
}
