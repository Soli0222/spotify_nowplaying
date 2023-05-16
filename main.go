package main

import (
	"net/http"
	"spotify_nowplaying/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/home", handlers.HomeHandler)
	http.HandleFunc("/callback", handlers.CallbackHandler)
	http.ListenAndServe(":8080", nil)
}
