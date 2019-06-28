package main

import (
    "time"
    "os"
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    time.Sleep(10 * time.Second)
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, world")
    })

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, os.Getenv("VERSION"))
    })
    e.GET("/api", func(c echo.Context) error {
        return c.String(http.StatusOK, "api")
    })

    e.Logger.Fatal(e.Start(":1323"))
}
