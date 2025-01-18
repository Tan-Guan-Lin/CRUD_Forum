package models

import (
	"time"
)

type UserResponse struct {
	Username string
	Password string
}

type User struct {
	UserResponse // Embedding userResponse struct
	UID          int64
}

type thread struct {
	TID           int64
	Username      string
	Title         string
	Content       string
	date_and_time time.Time
}

type comments struct {
	CID       int64
	Username  string
	Thread_ID int64
	Content   string
}
