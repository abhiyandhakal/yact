package main

import (
	"fmt"
	"net"
)

func main() {
	socketPath := "/tmp/yactd.sock"

	addr, err := net.ResolveUnixAddr("unix", socketPath)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to Unix socket:", socketPath)

	// Write and read data over the connection
	_, err = conn.Write([]byte("Hello from client!"))
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Printf("Received: %s\n", string(buf[:n]))
}
