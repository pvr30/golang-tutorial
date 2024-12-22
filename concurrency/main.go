package main

import (
	"fmt"
	"time"
)

// greet is a simple function that prints a message and signals completion via the channel.
func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true // Send a signal that the task is done.
}

// slowGreet simulates a long task by sleeping for 3 seconds before signaling completion.
func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // Simulate delay
	fmt.Println("Hello!", phrase)
	doneChan <- true // Send a signal that the task is done.
	close(doneChan)  // Close the channel (no more messages will be sent).
}

func main() {
	done := make(chan bool) // A channel for communication between goroutines.

	// Launch multiple goroutines, lightweight threads managed by Go runtime.
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

	// Use a range loop to read from the channel until it's closed.
	for range done {
		// Each message received indicates a completed task.
	}
}
