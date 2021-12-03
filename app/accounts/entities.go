package accounts

import "time"

type AccountEntity struct {
	Username  string
	Password  string
	CreatedAt time.Time
}
