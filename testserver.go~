package bustrack

import (
	"net/http"

	"strconv"

	"github.com/labstack/echo"
)

var users map[int]string

func main() {
	//create a new instance of echo
	e := echo.New()
	//define routes here fith anonymus functions or user defined functions
	//#1 with anonymus functions
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	//#2 with user defined functions
	//e.GET("/show", show)
	e.POST("/add", saveUser)
	e.GET("/get/:id", getUser)
	e.PUT("/update/", updateUser)
	e.DELETE("/delete/:id", deleteUser)
	e.Logger.Fatal(e.Start(":1321"))
}

/*func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}*/

func saveUser(c echo.Context) error {
	// Get name and email

	name := c.QueryParam("id")
	email := c.QueryParam("name")
	users[convertInt(name)] = string(email)
	return c.String(http.StatusCreated, "done")
	/*users[convertInt(c.QueryParam("id"))] = string(c.QueryParam("name"))
	// to use form data use this c.FormValue("name")
	return c.String(http.StatusCreated, "aaa")*/
}

func getUser(c echo.Context) error {
	return c.String(http.StatusOK, users[convertInt(c.Param("id"))])
}

func updateUser(c echo.Context) error {
	users[convertInt(c.QueryParam("id"))] = string(c.QueryParam("name"))
	return c.String(http.StatusOK, "done")
}
func deleteUser(c echo.Context) error {
	users[convertInt(c.Param("id"))] = ""
	return c.String(http.StatusOK, "done")
}

func convertInt(val string) int {
	a, _ := strconv.Atoi(val)
	return a
}
