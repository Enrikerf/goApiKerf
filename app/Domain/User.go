package Domain

import "time"

type User struct {
	Id        uint32
	Nickname  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}


