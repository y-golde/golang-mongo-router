package HandleErrors

import "net/http"

func InternalServerError(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error: " + e.Error()))
}