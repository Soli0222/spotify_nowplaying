package handlers

import (
	"net/http"
	"net/url"
	"os"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	auth_url := "https://accounts.spotify.com/authorize"

	client_id := os.Getenv("SPOTIFY_CLIENT_ID")
	redirect_uri := "http://127.0.0.1:8080/callback"
	scope := "user-read-currently-playing user-read-playback-state"

	params := url.Values{}
	params.Add("client_id", client_id)
	params.Add("response_type", "code")
	params.Add("redirect_uri", redirect_uri)
	params.Add("scope", scope)

	auth_url += "?" + params.Encode()

	http.Redirect(w, r, auth_url, http.StatusFound)
}
