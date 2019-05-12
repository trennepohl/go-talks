package main

import (
	"net/http"
	"text/template"

	"github.com/labstack/echo"
	//go get github.com/labstack/echo/...
)

type Pessoa struct {
	Name string
}

func main() {
	_, err := template.New("index.html").ParseFiles("assets/pages/index.html")
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Start(":8080")

}
