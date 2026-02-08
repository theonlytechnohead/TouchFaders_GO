package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	running = true
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go handle_signal(signals)

	address := getAddress()

	connect(address)
	go heartbeat()

	byte := readByte()
	fmt.Println("Read a byte!", byte)

	disconnect()
}

func handle_signal(signals <-chan os.Signal) {
	<-signals
	running = false
	fmt.Println("SIGNALLED!")
	disconnect()
	os.Exit(1)
}

func getAddress() string {
	var address string
	for len(address) == 0 && running {
		fmt.Print("Enter console IP address: ")
		fmt.Scanln(&address)
	}
	return address
}

func connect(address string) {
	stop := make(chan bool)
	stopped := make(chan bool)

	var fqdn = address + ":50000"
	fmt.Print("Connecting to ", fqdn)

	go dotcrawl(stop, stopped)

	var err error
	connection, err = net.Dial("tcp", fqdn)

	stop <- true
	<-stopped

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	connected = true
}

func disconnect() {
	connected = false

	if connection == nil {
		return
	}

	// to figure out if the connection is already closed, do a 1-byte read
	one := make([]byte, 1)
	connection.SetReadDeadline(time.Now())
	if _, err := connection.Read(one); err == io.EOF {
		connection.Close()
		connection = nil
	}
}

func readByte() byte {
	byte, err := bufio.NewReader(connection).ReadByte()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return byte
}

func heartbeat() {
	for connected && running {
		connection.Write([]byte{0xf0, 0x43, 0x10, 0x3e, 0x19, 0x7f, 0xf7})
		time.Sleep(1 * time.Second)
	}
}
