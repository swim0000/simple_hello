package main

import (
	"fmt"
	"net"
	"testing"
)

func TestPort3000IsListening(t *testing.T) {
	hostPort := "localhost:3000"
	conn, err := net.Dial("tcp", hostPort)
	if err != nil {
		t.Fatalf("Could not connect to %s: %v", hostPort, err)
	}
	defer conn.Close()

	fmt.Println("Port 3000 is listening.")
}
