package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func concurrencyWithGoRoutine() {
	// Set the value of GOMAXPROCS.
	runtime.GOMAXPROCS(1)
	// Add a count of two, one for each goroutine.
	wg.Add(2)
	fmt.Println("Start Goroutines")
	//launch a goroutine with label "A"
	go printCounts("A")
	//launch a goroutine with label "B"
	go printCounts("B")
	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

func printCounts(label string) {
	// Schedule the call to WaitGroup's Done to tell we are done.
	defer wg.Done()
	// Randomly wait
	for count := 1; count <= 10; count++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Count: %d from %s\n", count, label)
	}
}

// When you run the program, you see the following output. It will vary each time because of the random
// wait during the program execution:
// Start Goroutines
// Waiting To Finish
// Count: 1 from A
// Count: 1 from B
// Count: 2 from B
// Count: 2 from A
// Count: 3 from B
// Count: 3 from A
// Count: 4 from A
// Count: 5 from A
// Count: 4 from B
// Count: 6 from A
// Count: 5 from B
// Count: 7 from A
// Count: 6 from B
// Count: 7 from B
// Count: 8 from A
// Count: 8 from B
// Count: 9 from B
// Count: 9 from A
// Count: 10 from A
// Count: 10 from B
// Terminating Program