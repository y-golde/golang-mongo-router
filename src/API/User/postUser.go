package User

import (
	"context"
	"encoding/json"
	"helloworld/src/API/Common"
	"helloworld/src/API/Common/HandleErrors"
	"helloworld/src/DB/Connections"
	"helloworld/src/Entities"
	"net/http"
	"time"
)

func postUser(w http.ResponseWriter, r *http.Request) {
	Common.AddJsonHeader(w)
	
	var user Entities.User
	json.NewDecoder(r.Body).Decode(&user)

	users := Connections.GetUsersCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := users.InsertOne(ctx,user)
	if err != nil {
		HandleErrors.InternalServerError(w,err)
	}

	json.NewEncoder(w).Encode(res)
}