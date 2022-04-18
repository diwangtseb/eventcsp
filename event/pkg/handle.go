package pkg

import (
	"context"
	"fmt"
)

type Handle struct {
}

func NewHandle() *Handle {
	return &Handle{}
}

func (h *Handle) Handle(ctx context.Context, msg Msg) {
	// do something
	fmt.Printf("start handle event msg: %v \n", msg)
}
