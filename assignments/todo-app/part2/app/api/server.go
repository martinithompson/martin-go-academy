package api

import "net/http"

type Item struct {
	Name string `json:"name"`
}

func main() {
	http
	http.ListenAndServe(":5000", nil)
}
