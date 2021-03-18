package main

import (
	"fmt"
	"net"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Welcome")

	conn, err := net.DialTCP("tcp4", nil, &net.TCPAddr{IP: []byte{8, 8, 8, 8}, Port: 53})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Conn addr: ", conn.LocalAddr(), "Type: ", reflect.TypeOf(conn.LocalAddr()))
	fmt.Println("IP:", strings.Split(conn.LocalAddr().String(), ":")[0])
	fmt.Println("Port:", strings.Split(conn.LocalAddr().String(), ":")[1])
	conn.Close()
}
