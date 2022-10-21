package main
import(
	"fmt"
	"net"
	"os"
)

const(
	SERVER_HOST="localhost"
	SERVER_PORT="9199"
	SERVER_MODE="tcp"
)

func main(){
	fmt.Println("Server running ...")
	server, err := net.Listen(SERVER_MODE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
            fmt.Println("Error listening:", err.Error())
            os.Exit(1)
        }
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
    fmt.Println("Waiting for client...")
    for {
                connection, err := server.Accept()
                if err != nil {
                        fmt.Println("Error accepting: ", err.Error())
                        os.Exit(1)
                }
                fmt.Println("client connected")
                go processClient(connection)
        }
}

func processClient(connection net.Conn) {
        buffer := make([]byte, 1024)
        mLen, err := connection.Read(buffer)
        if err != nil {
                fmt.Println("Error reading:", err.Error())
        }
        fmt.Println("Received: ", string(buffer[:mLen]))
        _, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
        connection.Close()
}