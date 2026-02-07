package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

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

	return connection
}

/* https://stackoverflow.com/a/75962269/8705144 */
func collectChanOne[T any](ch <-chan T) (T, bool) {
	select {
	case val, stillOpen := <-ch:
		return val, stillOpen
	default:
		var zeroT T
		return zeroT, false
	}
}

func dotcrawl(stop chan bool, stopped chan bool) {
	crawl := true
	for crawl {
		fmt.Print(".")
		time.Sleep(100 * time.Millisecond)
		val, open := collectChanOne(stop)
		crawl = (!val || !open)
	}
	fmt.Println()
	stopped <- true
}

func disconnect(connection net.Conn) {
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
	connection.Write([]byte{0xf0, 0x43, 0x10, 0x3e, 0x19, 0x7f, 0xf7})
	time.Sleep(1 * time.Second)
}
