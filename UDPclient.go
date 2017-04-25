package mainq

import (
	"fmt"
	"net"
	"time"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp4", "192.168.1.102:8085")
	CheckError(err)

	//LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	//CheckError(err)

	Conn, err := net.DialUDP("udp4", nil, ServerAddr)
	CheckError(err)

	defer Conn.Close()
	i := 0
	for i<10 {
		msg := "trip1|\"lat:lan\"=>\"100.2:153.s\""
		i++
		buf := []byte(msg)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
}
