package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	event "github.com/diwangtseb/eventcsp/event/pkg"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/event?charset=utf8mb4"
	buffer := 10
	e := event.NewEvent(dsn, buffer, nil)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		for i := 0; i < 10; i++ {
			e.Collect(ctx, event.Msg{
				Name:    strconv.Itoa(i),
				Content: "hello world",
			})
			time.Sleep(time.Second * 1)
		}
	}()
	go e.StartProcess()
	time.Sleep(time.Second * 5)
	err := errors.New("s")
	if err != nil {
		cancel()
	}
	fmt.Println("end")
	time.Sleep(time.Second * 5)
}
