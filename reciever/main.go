package main

import (
	"bufio"
	"log"
	"net"
)

func main() {

	//run the connection
	connect(":8558")

}

//Connect to the server and listen for incomming traffic
func connect(port string) {

	//Try to form a connection
	sock, err := net.Listen("tcp", port)

	//Check for errors
	if err != nil {
		log.Print(err)
		return
	}

	// Close the connection when the function exits
	defer sock.Close()

	//Have an open connection and thread the listner
	for {
		conn, err := sock.Accept()

		//Check for errors
		if err != nil {
			log.Print(err)
			return
		}

		//handle the incoming requests
		go handleRequest(conn)
	}

}

// Handle the incomming requests
func handleRequest(conn net.Conn) {

	//Read any incomming data
	for {
		// Declare and initialise a buffer which will read from the
		// socket
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Print(err)
			return
		}
		log.Print(netData)
		break
	}
	conn.Close()
}
