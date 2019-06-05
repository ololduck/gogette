package main

import (
	"fmt"
	gnet "gogette/net"
	"net"
	"os"
	"strconv"
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
		request, err := gnet.CreateRequestFromBytes(buffer)
		var response *gnet.Response
		if err != nil {
			// Then, create a bad request response
			response = gnet.NewResponse(*gnet.BAD_REQUEST)
			response.Headers["Content-Type"] = "text/plain"
			response.Content = append(response.Content, []uint8("You sent a request I was unable to understand")...)
			response.Headers["Content-Length"] = strconv.Itoa(len(request.Content))
		} else {
			fmt.Printf("Generated Request Object: %s\n", request)
			response = gnet.NewResponse(*gnet.OK)
			response.Headers["Content-Type"] = "text/plain" // we just send the request back.
			response.Headers["Content-Length"] = strconv.Itoa(len(request.Content))
			response.Content = append(response.Content, []uint8("Here's what you sent: \n")...)
			response.Content = append(response.Content, request.Content...)
		}
		fmt.Printf("Response object: %s", *response)
		_, err = conn.Write(*response.ToBytes())
		if err != nil {
			fmt.Printf("Got an issue while writing: %s\n", err)
			break
		}
	}
	fmt.Println("Disconnect")
}
