package main

import (
	"fmt"
	"time"
)

/* https://stackoverflow.com/a/75962269/8705144 */
func readChannel[T any](channel <-chan T) (val T) {
	select {
	case value, _ := <-channel:
		return value
	default:
		var zeroT T
		return zeroT
	}
}

func dotcrawl(stop <-chan bool, stopped chan<- bool) {
	for !readChannel(stop) && running {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println()
	stopped <- true
}
