package main

import (
	"fmt"

	"github.com/rogue7373/GoLangScripts/blob/main/sampleProjects/CSTechTroubleShootingTool/hostname.go"
	"github.com/rogue7373/GoLangScripts/blob/main/sampleProjects/CSTechTroubleShootingTool/ipaddress.go"
	"github.com/rogue7373/GoLangScripts/blob/main/sampleProjects/CSTechTroubleShootingTool/username.go"
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
