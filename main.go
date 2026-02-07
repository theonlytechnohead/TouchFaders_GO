package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	var address string

	fmt.Print("Enter console IP address: ")
	fmt.Scan(&address)
	fmt.Println("The IP address is:", address)

	var fqdn = address + ":50000"

	conn, err := net.Dial("tcp", fqdn)
	if err != nil {
		fmt.Println(err)
		return
	}

	status, err := bufio.NewReader(conn).ReadByte()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Read a byte!", status)

	conn.Close()
}
