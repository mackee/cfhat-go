package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mackee/go-sqlla/v2"
	"github.com/syumai/workers"
	"github.com/syumai/workers/cloudflare/d1"
)

//go:embed static
var static embed.FS

func main() {
	router := httprouter.New()
	router.POST("/message", wrapHandler(func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) error {
		ctx := req.Context()
		db, err := connectDB(ctx)
		if err != nil {
			return fmt.Errorf("error connectDB: %w", err)
		}

		message := req.FormValue("message")
		if _, err := NewChatMessageSQL().Insert().
			ValueMessage(message).
			ExecContextWithoutSelect(ctx, db); err != nil {
			return fmt.Errorf("error Insert chatJmessage: %w", err)
		}

		http.Redirect(w, req, "/", http.StatusFound)
		return nil
	}))

	router.GET("/", wrapHandler(func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) error {
		tmpl := template.Must(template.ParseFS(static, "static/index.html"))

		ctx := req.Context()
		db, err := connectDB(ctx)
		if err != nil {
			return fmt.Errorf("error connectDB: %w", err)
		}

		messages, err := NewChatMessageSQL().Select().
			OrderByCreatedAt(sqlla.Desc).Limit(10).AllContext(ctx, db)
		if err != nil {
			return fmt.Errorf("error Select chat_message: %w", err)
		}
		type Message struct {
			Message   string
			CreatedAt string
		}
		ms := make([]Message, len(messages))
		for i, m := range messages {
			ms[i] = Message{
				Message:   m.Message,
				CreatedAt: m.CreatedAtStr(),
			}
		}

		if err := tmpl.Execute(w, ms); err != nil {
			return fmt.Errorf("error tmpl.Execute: %w", err)
		}

		return nil
	}))

	workers.Serve(router)
}

func wrapHandler(
	fn func(w http.ResponseWriter, req *http.Request, params httprouter.Params) error,
) func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	return func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
		if err := fn(w, req, params); err != nil {
			log.Printf("error %s", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func connectDB(ctx context.Context) (*sql.DB, error) {
	conn, err := d1.OpenConnector(ctx, "ChatDB")
	if err != nil {
		return nil, fmt.Errorf("error d1.OpenConnector: %w", err)
	}
	db := sql.OpenDB(conn)
	return db, nil

}
