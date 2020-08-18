package main

import (
	"fmt"
	"time"
)

// A defer statement defers the execution of a function until the surrounding function returns.
func defer_basic(){
	defer fmt.Println("World")
	fmt.Println("Hello")
}

// Deferred function calls are pushed onto a stack. When a function returns,
// its deferred calls are executed in last-in-first-out order.
func defer_stacked(){
	for i:=0; i < 10; i++ {
		defer fmt.Println("i: , time: ", i, time.Now())
	}
}

func main(){
	defer_basic()
	defer_stacked()
}