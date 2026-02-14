package main

import (
	"fmt"
	"time"
)

func BChannelsEx() {
	// Create a buffered channel with capacity 2
	ch := make(chan int, 2)

	// Sender goroutine
	go func() {
		fmt.Println("Sending message 1")
		ch <- 1
		fmt.Println("Sending message 2")
		ch <- 2
		fmt.Println("Attempting to send message 3 (will wait)")
		ch <- 3 // Blocks because buffer is full
		fmt.Println("Message 3 sent")
	}()

	// Simulate receiver delay
	time.Sleep(3 * time.Second)

	// Receiver
	fmt.Println("Receiving message:", <-ch)
	fmt.Println("Receiving message:", <-ch)
	fmt.Println("Receiving message:", <-ch)
}