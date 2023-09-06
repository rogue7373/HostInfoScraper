package main

import (
	"fmt"

	"github.com/rogue7373/hostname"
	"github.com/rogue7373/ipaddress"
	"github.com/rogue7373/username"
)

func main() {
	hostname, err := hostname.GetHostname()
	if err != nil {
		panic(err)
	}

	username, err := username.GetUsername()
	if err != nil {
		panic(err)
	}

	ipAddresses, err := ipaddress.GetIPAddress()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("IP Addresses: %s\n", ipAddresses)
}
