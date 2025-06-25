package user

import "time"

type User struct {
	UserID       string    `db:"user_id"`
	Name         string    `db:"name"`
	Image        *string   `db:"image"`
	PasswordHash string    `db:"password_hash"`
	PinHash      string    `db:"pin_hash"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type UserGreeting struct {
	UserID    string    `db:"user_id"`
	Greeting  string    `db:"greeting"`
	CreatedAt time.Time `db:"created_at"`
}
