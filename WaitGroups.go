package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, delay time.Duration) {
	defer wg.Done()

	fmt.Printf("Worker %d: Starting\n", id)
	time.Sleep(delay) 
	fmt.Printf("Worker %d: Completed\n", id)
}

func WaitGroupsEx() {
	var wg sync.WaitGroup

	numWorkers := 3

	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg, time.Duration(i)*time.Second)
	}

	wg.Wait()

	fmt.Println("All workers completed. Main process exits.")
}