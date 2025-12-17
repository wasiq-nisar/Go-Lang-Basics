package main

import (
	"fmt"
	"sync"
)

// wg is used to wait for the program to finish.
var wg2 sync.WaitGroup

func unBufferedChannels() {
	count2 := make(chan int)
	// Add a count of two, one for each goroutine.
	wg2.Add(2)
	fmt.Println("Start Goroutines")
	//launch a goroutine with label "A"
	go printCounts2("A", count2)
	//launch a goroutine with label "B"
	go printCounts2("B", count2)
	fmt.Println("Channel begin")
	count2 <- 1
	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg2.Wait()
	fmt.Println("\nTerminating Program")
}

func printCounts2(label string, count chan int) {
	// Schedule the call to WaitGroup's Done to tell we are done.
	defer wg2.Done()
	for {
	//Receives message from Channel
		val, ok := <-count
		if !ok {
		fmt.Println("Channel was closed")
		return
		}
		fmt.Printf("Count: %d received from %s \n", val, label)
		if val == 10 {
		fmt.Printf("Channel Closed from %s \n", label)
		// Close the channel
		close(count)
		return
	}
	val++
	// Send count back to the other goroutine.
	count <- val
	}
}

// The output can vary when you run the program. You should see output similar to the following:
// Start Goroutines
// Channel begin
// Count: 1 received from A
// Count: 2 received from B
// Waiting To Finish
// Count: 3 received from A
// Count: 4 received from B
// Count: 5 received from A
// Count: 6 received from B
// Count: 7 received from A
// Count: 8 received from B
// Count: 9 received from A
// Count: 10 received from B
// Channel Closed from B
// Channel was closed
// Terminating Program