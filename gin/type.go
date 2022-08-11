package main

import (
	"time"
)

type Data struct {
	Id        string    `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
