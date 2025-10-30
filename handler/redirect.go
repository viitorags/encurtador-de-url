package handler

import "net/http"

func Redirect(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		return
	}

	hash := r.URL.Path[1:]

	var originalURL string
	err := db.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", hash).Scan(&originalURL)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
