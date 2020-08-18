package main

import (
	"fmt"
)

type Vertex struct {
	X int;
	Y int;
}

func struct_basic(){
	obj := Vertex{4,5}
	fmt.Println(obj.X)
	fmt.Println(obj.Y)
}

func ptr_basic(){
	i,j:= 10, 50
	ptr := &i
	fmt.Println(*ptr)

	// This is known as "dereferencing" or "indirecting".
	*ptr = 30
	fmt.Println(*ptr)

	ptr = &j
	fmt.Println(*ptr)

}

func struct_ptr(){
	v:= Vertex{1,2}
	p:= &v
	p.X = 1000

	fmt.Println(v.X)
	fmt.Println(v.Y)

}

func main(){
	ptr_basic()
	fmt.Println("--------------------")
	struct_basic()
	fmt.Println("--------------------")
	struct_ptr()
}