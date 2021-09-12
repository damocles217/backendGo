package user

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/damocles217/server/models"
	"github.com/damocles217/server/router/user/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakeUser(
	w http.ResponseWriter,
	r *http.Request,
	collection *mongo.Collection) {

	// Set headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("limit", "50mb")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Define all variables
	var userReq models.UserRequest   // User catch request and get response
	var user models.User             // User sends to mongodb
	var errors []models.Error        // Errors matcher
	var setError models.Error        // Error type
	var response models.ResponseUser // Response

	// Set the user to request body
	err := json.NewDecoder(r.Body).Decode(&userReq)

	if err != nil {
		setError.Message = "Por favor envie datos validos"
		setError.Value = "[Body]"
		errors = append(errors, setError)
	}

	// Setting some values on user for mongodb
	user.Name = userReq.Name
	user.Lastname = userReq.Lastname
	user.Email = userReq.Email
	user.BornOn = userReq.BornOn
	user.Password, _ = config.HashPassword(userReq.Password)
	user.Logged = true
	user.Admin = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.CodeAuth = config.CodeAuth(15)
	user.UserID = config.GetUserID(user.Name, user.Lastname, collection)

	// Set cookie auth
	// crypting the code

	key := []byte("0123456789abcdef")
	result, err := config.AesEncrypt([]byte(user.CodeAuth), key)

	if err != nil {
		panic(err)
	}
	cryptedCode := base64.StdEncoding.EncodeToString(result)

	config.SetCookie(&w, "c_user", string(cryptedCode))

	// Search for user
	newUserId, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		setError.Message = "Internal server error!"
		setError.Value = "[Database]"

		errors = append(errors, setError)
	}

	userId, _ := newUserId.InsertedID.(primitive.ObjectID)

	filter := bson.D{primitive.E{Key: "_id", Value: userId}}

	collection.FindOne(context.TODO(), filter).Decode(&response.User)

	response.Errors = &errors

	if len(*response.Errors) > 0 {
		response.Sucess = false
	} else {
		response.Sucess = true
	}

	json.NewEncoder(w).Encode(response)

}
