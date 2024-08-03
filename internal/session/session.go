package session

import (
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

var Store *sessions.CookieStore

func Init() {
	secret := viper.GetString("server.secret")
	Store = sessions.NewCookieStore([]byte(secret))
}
