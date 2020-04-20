package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/aTTiny73/multilogger/logs"
)

func sendTime(conn *net.UDPConn, addreses *map[string]*net.UDPAddr, log *logs.MultipleLog) {

	for {

		time.Sleep(5 * time.Second)

		TIME := fmt.Sprint(time.Now().Format("15:04:05"))

		for _, v := range *addreses {

			data := []byte(TIME)

			_, err := conn.WriteToUDP(data, v)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func main() {

	serverlog := logs.NewFileLogger("Serverlog")
	defer serverlog.Close()

	stdlog := logs.NewStdLogger()
	defer stdlog.Close()

	log := logs.NewCustomLogger(false, serverlog, stdlog)

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	PORT := ":" + arguments[1]

	s, err := net.ResolveUDPAddr("udp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Listening port: %s \n", PORT)
	defer connection.Close()

	buffer := make([]byte, 1024)

	adresses := make(map[string]*net.UDPAddr)

	go sendTime(connection, &adresses, log)

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		log.Info(addr.String() + " says: " + string(buffer[0:n-1]))
		adresses[addr.String()] = addr

		data := []byte("Message recived")
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
