package main

import (
	"fmt"
	"time"
)

func main() {
	hourOfDay := time.Now().Hour()
	fmt.Println(getGreeting(hourOfDay))
}

func getGreeting(hour int) string {
	if hour < 12 {
		return "Good morning!"
	} else if hour < 18 {
		return "Good afternoon"
	} else {
		return "Good evening!"
	}
}
