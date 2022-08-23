package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
)

func Read(conn *pgx.Conn) (Data, error) {
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return Data{}, err
	}
	defer tx.Rollback(context.Background())

	sql := `SELECT id, content, created_at FROM todo`
	rows, err := tx.Query(context.Background(), sql)
	if err != nil {
		return Data{}, err
	}

	data := Data{}
	for rows.Next() {
		rows.Scan(
			&data.Id,
			&data.Content,
			&data.CreatedAt,
		)
	}

	return data, nil
}

func Insert(conn *pgx.Conn, content string) error {
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	sql := `INSERT INTO todo (content, created_at) VALUES ($1, $2)`
	tag, err := tx.Exec(context.Background(), sql, content, time.Now().Format(time.RFC3339))
	if err != nil {
		return err
	}
	fmt.Println(tag.RowsAffected(), tag.Insert())
	tx.Commit(context.Background())

	return nil
}

func Delete(conn *pgx.Conn, id string) error {
	sql := `DELETE FROM todo WHERE id = $1`
	_, err := conn.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}
	return nil
}
