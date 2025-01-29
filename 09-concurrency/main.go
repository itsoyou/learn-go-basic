package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan) // close the channel so that we does not get the error.
	// this only works when you know which function is the slowest.
}

func main() {

	/* 1. execution with goroutine.
	Program will end before they finishes */

	// go greet("Nice to meet you!")
	// go greet("How are you?")
	// go slowGreet("How ... are ... you ...?")
	// go greet("I hope you're liking the course!")

	/* 2. execution with goroutine and single channel.
	Program will end after the function execution. */

	// done := make(chan bool)
	// go slowGreet("How ... are ... you ...?", done)
	// <-done

	/* 3. Program will return different result every execution.
	We're using one channel for multiple goroutines.
	As soon as one function returns done value, program will exit */

	// done := make(chan bool)
	// go greet("Nice to meet you!", done)
	// go greet("How are you?", done)
	// go slowGreet("How ... are ... you ...?", done)
	// go greet("I hope you're liking the course!", done)
	// <-done

	/* 4. Program will wait until all functions return done value. */

	// done := make(chan bool)
	// go greet("Nice to meet you!", done)
	// go greet("How are you?", done)
	// go slowGreet("How ... are ... you ...?", done)
	// go greet("I hope you're liking the course!", done)
	// <-done
	// <-done
	// <-done
	// <-done

	/* 5. managing channel with slices. not ideal */

	// dones := make([]chan bool, 4)
	// dones[0] = make(chan bool)
	// go greet("Nice to meet you!", dones[0])
	// dones[1] = make(chan bool)
	// go greet("How are you?", dones[1])
	// dones[2] = make(chan bool)
	// go slowGreet("How ... are ... you ...?", dones[2])
	// dones[3] = make(chan bool)
	// go greet("I hope you're liking the course!", dones[3])
	// for _, done := range dones {
	// 	<-done
	// }

	/* 6. closing goroutine */

	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

	// for doneChan := range done {
	// 	fmt.Println(doneChan)
	// 	// fatal error: all goroutines are asleep - deadlock!
	// 	// because go doesn't know when this channel is out of values.
	// }

	for range done {
		// we closed the channel in the slowest function. via `close(doneChan)`
	}
}
