package handlers

import (
	"encoding/json"
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

	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+accessToken)
	headers.Set("ccept-Language", "ja")

	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/me", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header = headers

	response, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		http.Error(w, response.Status+" - user_data", http.StatusInternalServerError)
		return
	}

	var userData map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, userData["display_name"].(string))
}
