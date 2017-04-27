package controller

import (
	"bustrack/myredis"
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// Stream is a function to stream real time responses
func Stream(c echo.Context) error {
	res := c.Response()
	gone := res.CloseNotify()
	res.Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	res.WriteHeader(http.StatusOK)
	client := myredis.GetConnection()

	ticker := time.NewTicker(1 * time.Second)

	defer ticker.Stop()
	defer client.Close()

	client.Send("SUBSCRIBE", GetTopic(c.Param("tripid")))
	client.Flush()
	fmt.Fprint(res, "<pre><strong>Clock Stream</strong>\n\n<code>")

	for {
		result, _ := client.Receive()
		if result != nil {
			fmt.Fprintf(res, "%s\n", result)
		}
		res.Flush()
		select {
		case <-ticker.C:
		case <-gone:
			break
		}
	}
}

//GetTopic is method to make topic string to SUBSCRIBE
func GetTopic(tripid string) string {
	var buffer bytes.Buffer
	buffer.WriteString("trip")
	buffer.WriteString(tripid)
	return buffer.String()
}
