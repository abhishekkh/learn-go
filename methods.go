package main

import (
	"fmt"
)

type Vertex struct{
	X int
	Y int
}

// Method with type receiver
func (v Vertex) square() int {
	return v.X * v.X + v.Y + v.Y
}

// Method with pointer receiver
func (v *Vertex) scale(f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main(){
	v := Vertex{5, 2}
	fmt.Println("sqr: ", v.square())
	v.scale(4)
	fmt.Println("scaled: ", v.X, v.Y)

}