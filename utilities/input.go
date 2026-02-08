package utilities

import (
	"fmt"

	"github.com/theonlytechnohead/TouchFaders_GO/state"
)

func GetAddress() string {
	var address string
	for len(address) == 0 && state.Running {
		fmt.Print("Enter console IP address: ")
		fmt.Scanln(&address)
	}
	return address
}
