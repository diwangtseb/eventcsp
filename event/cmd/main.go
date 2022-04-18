package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	event "github.com/diwangtseb/eventcsp/event/pkg"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/event?charset=utf8mb4"
	buffer := 10
	e := event.NewEvent(dsn, buffer, nil)

	go func() {
		for i := 0; i < 10; i++ {
			e.Collect(context.Background(), event.Msg{
				Name:    strconv.Itoa(i),
				Content: "hello world",
			})
		}
	}()
	go e.StartProcess()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("exit", <-ch)
}
