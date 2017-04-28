package main

import (
	"bustrack/dbs"
	"bustrack/models"
	"bustrack/myredis"
	"bustrack/routes"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	poolCreatedStatus := myredis.InitPool()
	if poolCreatedStatus == true {
		e := echo.New()
		// Middleware
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		//CORS
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}))

		// Routes
		routes.Route(e)

		// Server
		go StartUdpServer()
		go func() {
			if err := e.Start(":1323"); err != nil {
				e.Logger.Info("shutting down the server")
			}
		}()
		// Wait for interrupt signal to gracefully shutdown the server with
		// a ttimeout of 10 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}
	//fmt.Println("initialization pool status:" + poolCreatedStatus)
}

// ------------------udp server code block starts here---------------------//
func handleUDPConnection(conn *net.UDPConn) {

	// here is where you want to do stuff like read or write to client

	buffer := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buffer)

	fmt.Println("UDP client : ", addr)
	fmt.Println("Received from UDP client :  ", string(buffer[:n]))
	UpdateTripCoords(string(buffer[:n]))

	if err != nil {
		log.Fatal(err)
	}

	// NOTE : Need to specify client address in WriteToUDP() function
	//        otherwise, you will get this error message
	//        write udp : write: destination address required if you use Write() function instead of WriteToUDP()

	if err != nil {
		log.Println(err)
	}

}

func StartUdpServer() {
	hostName := GetCurretIp()
	portNum := "8085"
	service := hostName + ":" + portNum

	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	fmt.Println(udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	// setup listener for incoming UDP connection
	ln, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UDP server up and listening on port 8085")

	defer ln.Close()

	for {
		// wait for UDP client to connect
		handleUDPConnection(ln)
	}

}

func GetCurretIp() string {
	addrs, err := net.InterfaceAddrs()
	currentIP := "0.0.0.0"
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				currentIP = ipnet.IP.String()
			}
		}
	}
	return currentIP
}

//------------------------------udp server block ends here-----------------//
//UpdateTripCoords is to add current location to existing trip
func UpdateTripCoords(locationstamp string) {
	var coords, topic, tripid string
	temp := strings.Split(locationstamp, "|")
	topic = temp[0]
	coords = temp[1]

	//to publish to the redis topic
	go publishToRedis(topic, coords)
	tripid = strings.SplitAfter(topic, "trip")[1]

	//to update things in database
	go func() {
		_, err := models.UpdateTripDetails(dbs.DB, coords, stringtoInt(tripid))

		if err != nil {
			fmt.Println("Error is", err)
		}
	}()

}

func stringtoInt(s string) int {
	str, _ := strconv.Atoi(s)
	return str
}

func publishToRedis(topic, coords string) {
	client := myredis.GetConnection()
	defer client.Close()
	fmt.Println(coords)
	_, err := client.Do("PUBLISH", topic, coords)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
	}
}
