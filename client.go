package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func readUDP(conn *net.UDPConn) {
	for {
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Server: %s\n", string(buffer[0:n]))
	}
}

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string")
		return
	}
	CONNECT := arguments[1]

	addr, err := net.ResolveUDPAddr("udp", CONNECT)
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s\n", conn.RemoteAddr().String())
	defer conn.Close()

	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text)
		_, err = conn.Write(data)

		if err != nil {
			fmt.Println(err)
			return
		}
		go readUDP(conn)
	}
}
