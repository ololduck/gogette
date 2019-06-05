package main

import (
	"fmt"
	gnet "gogette/net"
	"net"
	"os"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Listen error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Now listening on 8080")
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept error: %s\n", err)
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	fmt.Println("Got a connection")
	defer conn.Close()
	//NICE: A bufio.Scanner custom implementation
	buffer := make([]uint8, 4096)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Got an issue while reading: %s\n", err)
			break
		}
		request := gnet.CreateRequestFromBytes(buffer)
		fmt.Printf("Generated Request Object: %s\n", request)
		response := gnet.NewResponse(*gnet.OK)
		response.Headers["Content-Type"] = "text/plain" // we just send the request back.
		response.Content = append(response.Content, []uint8("Here's what you sent: \n")...)
		response.Content = append(response.Content, request.Content...)
		_, err = conn.Write(response.ToBytes())
		if err != nil {
			fmt.Printf("Got an issue while writing: %s\n", err)
			break
		}
	}
	fmt.Println("Disconnect")
}
