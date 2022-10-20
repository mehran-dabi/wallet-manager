package entity

import "time"

type Wallet struct {
	ID        int64     `json:"id"`
	Balance   uint64    `json:"balance"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}
