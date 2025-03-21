package network

import (
	"fmt"
	"net"
)

// Start P2P Server
func StartServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting P2P server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("P2P Server listening on port", port)
	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection:", conn.RemoteAddr().String())
}
