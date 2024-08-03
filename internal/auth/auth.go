package auth

import "github.com/gorilla/sessions"

type Users struct {
	store sessions.Store
}
