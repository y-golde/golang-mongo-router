package Users

import (
	"context"
	"encoding/json"
	"helloworld/src/API/Common"
	"helloworld/src/API/Common/HandleErrors"
	"helloworld/src/DB/Connections"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)


func getUsers(w http.ResponseWriter, r *http.Request) {
	Common.AddJsonHeader(w)
	
	usersCollection := Connections.GetUsersCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	cur, err := usersCollection.Find(ctx, bson.M{})
	if err != nil {
		defer cur.Close(ctx)
		HandleErrors.InternalServerError(w,err)
	}

	var users []bson.M
	err = cur.All(ctx, &users)
	if err != nil { HandleErrors.InternalServerError(w,err) }

	json.NewEncoder(w).Encode(users)
}