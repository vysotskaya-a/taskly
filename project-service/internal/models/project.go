package models

import "time"

type Project struct {
	ID                           string    `json:"id" db:"id"`
	Title                        string    `json:"title" db:"title"`
	Description                  string    `json:"description" db:"description"`
	Users                        []string  `json:"users" db:"users"`
	AdminID                      string    `json:"admin_id" db:"admin_id"`
	NotificationSubscribersTGIDS []int64   `json:"notification_subscribers_tg_ids" db:"notification_subscribers_tg_ids"`
	CreatedAt                    time.Time `json:"created_at" db:"created_at"`
}
