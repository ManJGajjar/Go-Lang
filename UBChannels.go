package main

import (
	"fmt"
	"time"
)

func UBChannelsEx() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello, Gophers!"
	}()

	fmt.Println("Waiting for sender...")
	message := <-ch 
	fmt.Println("Received:", message)
}