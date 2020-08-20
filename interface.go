package main

import(
	"fmt"
)

type emitter interface {
	//this method is implemented
	emit(topic string, message string)
}

type MyJsonEmitter struct {
	size int
	messageType string
}

// notice no pointer receiver
func (e *MyJsonEmitter) emit(topic string, message string){
	fmt.Printf("%v, %T\n", e, e)
	fmt.Println("emitting message: ", message)
}

func main(){
	var e emitter = &MyJsonEmitter{5, "json"}
	e.emit("xyz", "Hello! How are you?")

	// with null values
	var e1 emitter
	var j1 *MyJsonEmitter
	e1 = j1
	e1.emit("xyz", "nil interface")
}