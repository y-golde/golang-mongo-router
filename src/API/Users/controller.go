package Users

import (
	"net/http"
)

func UsersController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET" :
			getUsers(w,r)
	}
}