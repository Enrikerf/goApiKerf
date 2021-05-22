package Domain

import "time"

type User struct {
	id        uint32
	nickname  string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}


