package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), URL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	runRouter(conn)
}
