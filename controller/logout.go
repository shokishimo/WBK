package controller

import (
	"context"
	"github.com/shokishimo/WhatsTheBestKeyboard/database"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := GetSessionCookie(r)

	// delete session handling cookie from the browser
	DeleteCookie(w, "username", "")
	DeleteCookie(w, "sessionid", "")

	// delete the session id of the user in the database
	hashedId := Hash(sessionId)
	db := database.Connect()
	defer db.Disconnect()
	db = db.GetAccessKeysToUsersCollection()

	filter := bson.M{"sessionid": hashedId}
	update := bson.M{"$set": bson.M{"sessionid": ""}}
	result, err := db.GetCollection().UpdateOne(context.TODO(), filter, update)

	// when an error happened in the transaction
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error happened during the transaction"))
		return
	}
	// when the user with the username and password not found
	if result.MatchedCount == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error: user not found"))
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
