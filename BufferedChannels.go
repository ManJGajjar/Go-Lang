package main

import (
	"fmt"
	"time"
)

func BChannelsEx() {
	ch := make(chan int, 2)

	go func() {
		fmt.Println("Sending message 1")
		ch <- 1
		fmt.Println("Sending message 2")
		ch <- 2
		fmt.Println("Attempting to send message 3 (will wait)")
		ch <- 3 
		fmt.Println("Message 3 sent")
	}()

	time.Sleep(3 * time.Second)

	fmt.Println("Receiving message:", <-ch)
	fmt.Println("Receiving message:", <-ch)
	fmt.Println("Receiving message:", <-ch)

}
