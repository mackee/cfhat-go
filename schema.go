package main

import (
	"github.com/mackee/go-genddl/index"
)

//go:generate go run github.com/mackee/go-sqlla/v2/cmd/sqlla
//go:generate go run github.com/mackee/go-genddl/cmd/genddl -outpath=./schema.sql -driver=sqlite3

type ChatMessageID int64

// +table: chat_message
type ChatMessage struct {
	ID        ChatMessageID `db:"id,primarykey,autoincrement"`
	Message   string        `db:"message,text"`
	CreatedAt int64         `db:"created_at"`
}

func (c ChatMessage) _schemaIndex(methods index.Methods) []index.Definition {
	return []index.Definition{
		methods.Complex(c.CreatedAt),
	}
}
