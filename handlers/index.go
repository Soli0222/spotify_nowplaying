package handlers

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	html := "<html><head><title>My Page</title></head><body><h1>Welcome to my page</h1></body></html>"
	w.Write([]byte(html))

}
