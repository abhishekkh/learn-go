package main

import(
	"fmt"
)

func main(){
	var sum int
	for i:=0; i<10; i++ {
		fmt.Println("Hello World")
		sum+=i
	}
	fmt.Println("sum: ", sum)

}