package network

import (
	"fmt"
	"net"
	"os"

	"github.com/theonlytechnohead/TouchFaders_GO/utilities"
)

func Connect(address string) (connection net.Conn, connected bool) {
	stop := make(chan bool)
	stopped := make(chan bool)

	var fqdn = address + ":50000"
	fmt.Print("Connecting to ", fqdn)

	go utilities.DotCrawl(stop, stopped)

	connection, err := net.Dial("tcp", fqdn)

	stop <- true
	<-stopped

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return connection, true
}
