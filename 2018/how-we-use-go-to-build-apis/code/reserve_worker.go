package main

import (
	"fmt"
	"time"
)
// START OMIT
func main() {
	jobs := make(chan chan int) // first chan to reserve a worker, second one to send a task // HL
	for i := 0; i < 2; i++ {
		go worker(jobs)
	}

	for i := 0; i < 10; i++ {
		job := make(chan int)
		select {
			case jobs <- job: // HL
				fmt.Println("reserve worker for task:", i)
				job <- i // build and send the task // HL
			default:
				fmt.Println("drop task:", i)
		}
		time.Sleep(time.Millisecond/2) // OMIT
	}
}

func worker(jobs chan chan int) {
	for job := range jobs {
		task := <- job
		time.Sleep(time.Millisecond) // OMIT
		fmt.Println("done task:", task)
	}
}
// END OMIT

