package controller

import (
	"bustrack/myredis"
	"bytes"
	"fmt"
	"log"

	"golang.org/x/net/websocket"

	"github.com/labstack/echo"
)

// Stream is a function to stream real time responses
func Stream(c echo.Context) error {

	websocket.Handler(func(ws *websocket.Conn) {
		client := myredis.GetConnection()
		defer client.Close()
		defer ws.Close()

		client.Send("SUBSCRIBE", GetTopic(c.Param("tripid")))
		client.Flush()

		for {
			// Write
			result, _ := client.Receive()
			if result != nil {
				err := websocket.Message.Send(ws, fmt.Sprintf("%s\n", result))
				if err != nil {
					log.Fatal(err)
				}
			}

		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

//GetTopic is method to make topic string to SUBSCRIBE
func GetTopic(tripid string) string {
	var buffer bytes.Buffer
	buffer.WriteString("trip")
	buffer.WriteString(tripid)
	return buffer.String()
}
