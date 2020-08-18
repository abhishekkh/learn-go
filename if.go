package main

import(
	"fmt"
)

func check_name (name string) bool {
	if name == "abhi"{
		return true
	}
	return false
}

func main(){
	name := "abhi"
	if check_name(name){
		fmt.Println("Valid name provided")	
	} 
}