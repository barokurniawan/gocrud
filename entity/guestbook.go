package entity

import "time"

type Guestbook struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Message   string    `json:"message"`
	CreatedAT time.Time `json:"created_at"`
}
