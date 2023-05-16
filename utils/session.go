package utils

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func init() {
	// セッションストアの初期化
	store = sessions.NewCookieStore([]byte("secret-key"))
}

func GetSession(r *http.Request) (*sessions.Session, error) {
	// セッションの取得
	session, err := store.Get(r, "session-name")
	if err != nil {
		return nil, err
	}

	return session, nil
}
