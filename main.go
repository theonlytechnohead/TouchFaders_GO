package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

var connected = make(chan bool)

func main() {
	address := getAddress()

	connection := connect(address)
	go heartbeat(connection)

	byte := readByte(connection)
	fmt.Println("Read a byte!", byte)

	disconnect(connection)
}

func getAddress() string {
	var address string
	for len(address) == 0 {
		fmt.Print("Enter console IP address: ")
		fmt.Scanln(&address)
	}
	return address
}

func connect(address string) net.Conn {
	stop := make(chan bool)
	stopped := make(chan bool)

	var fqdn = address + ":50000"
	fmt.Print("Connecting to ", fqdn)

	go dotcrawl(stop, stopped)

	connection, err := net.Dial("tcp", fqdn)

	stop <- true
	<-stopped

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	connected <- true

	return connection
}

/* https://stackoverflow.com/a/75962269/8705144 */
func readChannel[T any](channel <-chan T) (val T, open bool) {
	select {
	case value, open := <-channel:
		return value, open
	default:
		var zeroT T
		return zeroT, false
	}
}

func dotcrawl(stop <-chan bool, stopped chan<- bool) {
	crawl := true
	for crawl {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
		_stop, open := readChannel(stop)
		crawl = (!_stop || !open)
	}
	fmt.Println()
	stopped <- true
}

func disconnect(connection net.Conn) {
	connected <- false
	connection.Close()
}

func readByte(connection net.Conn) byte {
	byte, err := bufio.NewReader(connection).ReadByte()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return byte
}

func heartbeat(connection net.Conn) {
	beat := true
	for beat {
		connection.Write([]byte{0xf0, 0x43, 0x10, 0x3e, 0x19, 0x7f, 0xf7})
		time.Sleep(1 * time.Second)
		beat, _ = readChannel(connected)
	}
}
