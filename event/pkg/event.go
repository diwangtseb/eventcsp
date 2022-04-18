package pkg

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
)

type Event struct {
	db      *sql.DB        //mysql db
	log     *logrus.Logger //logger
	ch      chan *Msg      //event msg channel
	handler Handler        //handler
}

func NewEvent(dsn string, buffer int, handler Handler) *Event {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil
	}
	var h Handler
	if handler == nil {
		h = NewHandle()
	} else {
		h = handler
	}
	return &Event{
		db:      db,
		ch:      make(chan *Msg, buffer),
		handler: h,
		log:     logrus.New(),
	}
}

func (e *Event) Collect(ctx context.Context, msg Msg) {
	e.ch <- &msg
}

func (e *Event) StartProcess() {
	for msg := range e.ch {
		e.handler.Handle(context.Background(), *msg)
	}
}

/**
 * @description: close all connect
 * @param {*}
 * @return {*}
 */
func (e *Event) Close() {
	e.db.Close()
}
