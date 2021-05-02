package User

import (
	"net/http"
)

func UserController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST" :
			postUser(w,r)
	}
}