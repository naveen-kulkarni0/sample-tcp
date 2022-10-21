package main

import(
	"fmt"
	"net"
)

const (
	SERVER_HOST="localhost"
	SERVER_PORT="9199"
	SERVER_MODE="tcp"
)

func main(){
	//establish connection
	connection,err := net.Dial(SERVER_MODE, SERVER_HOST + ":" + SERVER_PORT)
	if(err != nil){
		panic(err)
	}

	_,err = connection.Write([]byte("Hello Server! Greetings."))
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	 if err != nil {
                fmt.Println("Error reading:", err.Error())
        }
	fmt.Println("Received: ", string(buffer[:mLen]))
    defer connection.Close()

}
