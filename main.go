package main

import (
	"fmt"
	"bhyveconnect"
)

func main() {
	fmt.Printf("Hello B-Hyve\n")

	//client := bhyveconnect.NewBasicAuthClient("me@campbellmcneill.com", "r3dd3ad@#")

	// Get a Session
	//s, _ := bhyveconnect.Login(&client)

	bhyveconnect.HelloWorld()
}
