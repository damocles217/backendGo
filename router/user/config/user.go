package config

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/damocles217/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Easy way to create a cookie
func SetCookie(w *http.ResponseWriter, name string, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Secure:   true,
		HttpOnly: true,
		MaxAge:   3600 * 24 * 60,
	}

	http.SetCookie(*w, cookie)
}

// Make an auth code for email
func CodeAuth(length int) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Check the key "userId" in the database to check the key'd be unique
func GetUserID(name string, lastname string, collection *mongo.Collection) string {
	var userId = name + "." + lastname
	var err error
	var checkUser models.User
	filter := bson.D{primitive.E{Key: "userId", Value: userId}}

	for i := 1; i > 0; i++ {

		err = collection.FindOne(context.Background(), filter).Decode(&checkUser)

		if err == nil {
			userId = name + "." + lastname + strconv.Itoa(i)
			filter = bson.D{primitive.E{Key: "userId", Value: userId}}
		} else {
			i = -1
		}
	}

	return userId
}
