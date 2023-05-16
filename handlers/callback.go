package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"spotify_nowplaying/utils"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token_url := "https://accounts.spotify.com/api/token"

	redirect_uri := "http://127.0.0.1:8080/callback"
	client_id := os.Getenv("SPOTIFY_CLIENT_ID")
	client_secret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("code", code)
	params.Add("redirect_uri", redirect_uri)
	params.Add("client_id", client_id)
	params.Add("client_secret", client_secret)

	response, err := http.PostForm(token_url, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accessToken, ok := data["access_token"].(string)
	if !ok {
		http.Error(w, "access_token not found", http.StatusInternalServerError)
		return
	}

	session, err := utils.GetSession(r)
	if err != nil {
		// セッションの取得に失敗した場合のエラーハンドリング
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["access_token"] = accessToken

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/home", http.StatusFound)
}
