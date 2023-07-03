package types

import "time"

type WebProcess struct {
	Name       string    `gorm:"primary_key;unique" json:"name"`
	Active     bool      `json:"active"`
	StartedAt  time.Time `json:"started_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeployedAt time.Time `json:"deployed_at"`
}

type WebError struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type WebProcessOrError struct {
	Process WebProcess `json:"process"`
	WebError
}
