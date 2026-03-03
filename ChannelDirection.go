package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
    for task := range tasks {
        result := processTask(task)
        results <- result
        fmt.Printf("Worker %d processed task %d\n", id, task)
    }
    wg.Done()
}

func processTask(task int) int {
    time.Sleep(time.Second)  
    return task * 2
}

func ChannelDirectionEx() {
    const numWorkers = 3
    const numTasks = 5

    tasks := make(chan int, numTasks) 
    results := make(chan int, numTasks) 
    var wg sync.WaitGroup

    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, tasks, results, &wg)
    }

    for j := 1; j <= numTasks; j++ {
        tasks <- j
    }
    close(tasks)
    wg.Wait()
    close(results) 

    for result := range results {
        fmt.Println("Result:", result)
    }

}
