package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	// we need to close the channels using defer at the end of the method
	// otherwise we get -> fatal error: all goroutines are asleep - deadlock!
	go func() {
		defer close(ch1)
		Walk(t1, ch1)
	}()

	go func() {
		defer close(ch2)
		Walk(t2, ch2)
	}()

	for {
		x, ok1 := <- ch1
		y, ok2 := <- ch2
		fmt.Println(x,y)
		
		if x!=y || ok1!=ok2 {
			return false
		}
		
		if !ok1 {
            break
        }
	}
	
	 for k := range ch1 {
        select {
        case g := <-ch2:
            if k != g {
                return false
            }
        default:
            break
        }
    }
	return true
}

func main() {
	t1 := tree.New(3)
	t2 := tree.New(3)
	fmt.Println(t1, t2)
	fmt.Printf("Result: %v",Same(t1, t2))
}
