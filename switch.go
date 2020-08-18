package main

import (
	"fmt"
	"time"
)

func switch_with_no_statement(){
	t:=time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")
	case t.Hour() < 18:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good night")
	}
}


func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	switch_with_no_statement()
}