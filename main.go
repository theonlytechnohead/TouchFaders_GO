package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/theonlytechnohead/TouchFaders_GO/network"
	"github.com/theonlytechnohead/TouchFaders_GO/state"
	"github.com/theonlytechnohead/TouchFaders_GO/utilities"
)

func main() {
	state.Running = true
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go handleSignal(signals)

	address := utilities.GetAddress()

	state.Connection, state.Connected = network.Connect(address)
	go heartbeat()

	byte := readByte()
	fmt.Println("Read a byte!", byte)

	network.Disconnect()
}

func handleSignal(signals <-chan os.Signal) {
	<-signals
	fmt.Println()
	fmt.Print("Exiting neatly...")
	state.Running = false
	network.Disconnect()
	fmt.Println(" DONE!")
	os.Exit(1)
}

func readByte() byte {
	byte, err := bufio.NewReader(state.Connection).ReadByte()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return byte
}

func heartbeat() {
	for state.Connected && state.Running {
		state.Connection.Write([]byte{0xf0, 0x43, 0x10, 0x3e, 0x19, 0x7f, 0xf7})
		time.Sleep(1 * time.Second)
	}
}
