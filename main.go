package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var count = new(int32)

func main() {
	e := echo.New()
	resources := make(chan interface{}, 10)
	go func(ch chan interface{}) {
		for true {
			// time.Sleep(time.Duration(500) * time.Millisecond)
			CumbersomeFunction()
			ch <- new(interface{})
		}
	}(resources)
	e.GET("/", getDemoRequestHandler(resources))
	e.GET("/health", handleHealthCheck)
	e.Logger.SetLevel(log.INFO)
	e.Logger.Fatal(
		e.Start(":8082"),
	)
}

func getDemoRequestHandler(resources chan interface{}) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		start := time.Now()
		<-resources
		ctx.Logger().Info(fmt.Sprintf("%d", atomic.AddInt32(count, 1)))
		return ctx.String(http.StatusOK, fmt.Sprintf("%v", time.Now().Sub(start).Milliseconds()))
	}
}

func handleHealthCheck(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

// CumbersomeFunction performs long calculations
func CumbersomeFunction() {
	value := []byte("Считаем от рассвета до заката")
	for i := 0; i < rand.Intn(1000000); i++ {
		hash := sha256.Sum256(value)
		value = append(value, hash[:]...)
	}
}
