package model

import "time"

type Item struct {
	ID        int
	Title     string
	Desc      string
	Price     int
	createdAt time.Time
	updatedAt time.Time
}

type Mst_user struct {
	Mus_id       int    `json:"mus_id"`
	Mus_email    string `json:"mus_email"`
	Mus_fullname string `json:"mus_fullname"`
	Mus_password string `json:"mus_password"`
	Mus_mac      string `json:"mus_mac"`
}
