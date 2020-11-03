package common

import (
	"fmt"
	"net"
	"strings"
)

//udp config
const (
	SERVER_RECV_LEN = 65000
)

//StartUDPServer : udp server
func StartUDPServer(port int, output chan string) {

	address := fmt.Sprintf("0.0.0.0:%d", port)

	fmt.Println("Starting UDP Server with address ", address)

	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	go func() {

		buff := make([]byte, SERVER_RECV_LEN)

		for {

			// Here must use make and give the length of buffer
			len, rAddr, err := conn.ReadFromUDP(buff)

			if err != nil {
				fmt.Println(err)
				continue
			}

			data := string(buff[:len])
			//fmt.Println("Received:", data)

			output <- data

			upper := strings.ToUpper(data)
			_, err = conn.WriteToUDP([]byte(upper), rAddr)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("Send:", upper)
		}
	}()

	select {}

}
