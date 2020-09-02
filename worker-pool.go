package main

import (
	"fmt"
	"time"
)

func worker (id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(1 * time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- 2 * j
	}
}


func main() {
	numJobs := 5
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w:=0; w < 3; w++ {
		go worker(w, jobs, results)
	}

	for i:=0; i < numJobs; i++ {
		jobs <- i
	}

	// close the job channel, let the receiver know there are no more jobs
	close(jobs)

	// wait for all the results to be completed
	for j:=0; j < numJobs; j++ {
		<- results
	}


}