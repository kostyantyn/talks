package main

import (
	"net/http"
)

// START OMIT
func main() {
	http.HandleFunc("/basic_auth", func(w http.ResponseWriter, r *http.Request) {
		usr, pwd, _ := r.BasicAuth()
		if usr != "user" || pwd != "password" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Application"`)
			w.WriteHeader(401)
		}
	})

	http.ListenAndServe(":8080", nil)
}
// END OMIT
