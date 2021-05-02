package API

import (
	"fmt"
	"helloworld/src/API/User"
	"helloworld/src/API/Users"
	"net/http"
	"os"
)

var PORT string = os.Getenv("API_PORT")

func InitRouter() {
	m := http.NewServeMux()

	m.HandleFunc("/users", Users.UsersController)
	m.HandleFunc("/user", User.UserController)
	m.HandleFunc("/helo", helo)

	s := &http.Server{
		Addr: ":1948",
		Handler: m,
	}
	
	fmt.Printf("server started on Port ");
	s.ListenAndServe()
}

func helo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("helo"))
	w.Write([]byte(os.Getenv("DB_URL")))
}