package handlers

import (
	"fmt"
	"net/http"
	"spotify_nowplaying/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, err := utils.GetSession(r)
	if err != nil {
		// セッションの取得に失敗した場合のエラーハンドリング
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accessToken, ok := session.Values["access_token"].(string)
	if !ok {
		http.Error(w, "access_token not found", http.StatusInternalServerError)
		return
	}

	// アクセストークンを使用して何らかの処理を行う
	fmt.Fprintf(w, accessToken)
}
