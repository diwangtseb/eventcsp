package pkg

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
)

type Msg struct {
	Name    string
	Contetn string
}

type Eventer interface {
	Collect(ctx context.Context, msg Msg) // collect events
}

type Handler interface {
	Handle(ctx context.Context, msg Msg) // handle events
}
