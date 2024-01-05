package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting \n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done \n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		fmt.Println("adding => ", i)
		wg.Add(1)
		counter := i
		// order of execution of the goroutines is not guaranteed
		// nor is time of execution for each routine
		go func() {
			defer wg.Done()
			worker(counter)
		}()
	}
	fmt.Println("End")
	wg.Wait()
	fmt.Println("After WG")
}
