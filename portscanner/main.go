package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	if len(os.Args) <= 3 {
		fmt.Print("Missing params (host, start port, end port)\n")
		os.Exit(1)
	}

	host := os.Args[1]

	start, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("First port number must be an integer")
		os.Exit(1)
	}

	end, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Second port number must be an integer")
		os.Exit(1)
	}
	portscanner(host, start, end)
}


func portscanner (address string, start, end int) {

	timeoutseconds := 2
	host := "tcp"

	fmt.Println("Connection is open on the following ports:")

	for s := start; s <= end; s++ {
		timeout := time.Duration(timeoutseconds) * time.Second
		newstart := strconv.Itoa(s)
		newAddress := strings.Join([]string{address, ":", newstart}, "")

		conn, err := net.DialTimeout(host, newAddress, timeout)
		if err != nil {
			//fmt.Printf("Failed on port %d\n", s)
		} else {
			fmt.Printf("Success on remote address: %s \n", conn.RemoteAddr().String())
			//fmt.Printf("Local address: %s \n", conn.LocalAddr().String())
		}
	}
	return
}


