package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {

	//Run the proxy
	Proxy()
}

// Proxy function which takes the request and will
// forward it on to the relevent server
func Proxy() {

	//will serve the http to where it needs to go
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//Check if there is an error here
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
			return
		}

		//Extract the string
		str := r.FormValue("String")
		str = str + "\n"

		send(str)

	})

	// Listen on :6285
	log.Fatal(http.ListenAndServe(":6285", nil))

}

//Connect to the server and send the body
func send(str string) {

	//Try to form a connection
	conn, err := net.Dial("tcp", "localhost:8558")

	//Check for errors
	if err != nil {
		fmt.Println(err)
		return
	}

	// Close the connection when the function exits
	defer conn.Close()

	//Check for errors
	if err != nil {
		fmt.Println(err)
		return
	}

	//Send the request to the server
	bte := []byte(str)
	_, err = conn.Write(bte)

	//Check for errors
	if err != nil {
		fmt.Println(err)
		return
	}
}
