package Common

import "net/http"

func AddJsonHeader(w http.ResponseWriter) {
	w.Header().Add("content-type", "application-json")
}