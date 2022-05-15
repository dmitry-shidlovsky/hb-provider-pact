package model

import "errors"

type User struct {
	FirstName string `json:"firstName" pact:"example=Dmitry"`
	LastName  string `json:"lastName" pact:"example=Shidlovsky"`
	Username  string `json:"username" pact:"example=dsh"`
	Type      string `json:"type" pact:"example=admin,regex=^(admin|user|guest)$"`
	ID        int    `json:"id" pact:"example=1"`
}

var (
	ErrNotFound = errors.New("not found")
)
