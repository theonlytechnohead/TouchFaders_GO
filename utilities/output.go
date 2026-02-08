package utilities

import (
	"fmt"
	"time"

	"github.com/theonlytechnohead/TouchFaders_GO/state"
)

func DotCrawl(stop <-chan bool, stopped chan<- bool) {
	for !ReadChannel(stop) && state.Running {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println()
	stopped <- true
}
