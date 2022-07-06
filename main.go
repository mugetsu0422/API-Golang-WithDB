package main

import(
	"net/http"
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
	"API-Golang-WithDB/storage"
	"API-Golang-WithDB/controller"
)

func main() {
	// Echo instance 
	e := echo.New()
	storage.NewDB()
	// Middleware
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/students", controller.GetAllStudents)
	e.GET("/students/:id", controller.GetStudent)
	e.POST("/students", controller.CreateStudent)
	e.PUT("/students/:id", controller.UpdateStudent)
	e.DELETE("/students/:id", controller.DeleteStudent)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World")
}
