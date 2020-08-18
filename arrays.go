package main

import (
	"fmt"
)

func splicing(){
	flags:=[]bool{true, false, false}
	fmt.Println(flags[1])

	points:=[]struct {
		x int
		y int
	}{
		{4, 6},
		{9, 10},
		{16, 22345},
	}
	fmt.Println(points)
}

func main(){
	var a[2]int
	a[0] = 1
	a[1] = 5
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	splicing()
}