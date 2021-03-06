package mainqw

import (
	"bustrack/myredis"
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
)

func stream(c echo.Context) error {
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

func main() {
	poolCreatedStatus := myredis.InitPool()
	if poolCreatedStatus == true {
		e := echo.New()
		e.GET("/stream/:tripid", stream)
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

	fmt.Println(poolCreatedStatus)

}

func GetTopic(tripid string) string {
	var buffer bytes.Buffer
	buffer.WriteString("trip")
	buffer.WriteString(tripid)
	return buffer.String()
}
