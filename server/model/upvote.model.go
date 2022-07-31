package model

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/alexsandron3/klever-test/server/service"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// TO-DO = Refact this code to use ENV vars and REMOTE database
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while trying to load .env file")
	}
	mongoUser := os.Getenv("mongoUser")
	mongoPass := os.Getenv("mongoPass")
	mongoCluster := os.Getenv("mongoCluster")
	uri := "mongodb+srv://" + url.QueryEscape(mongoUser) + ":" + url.QueryEscape(mongoPass) + "@" + mongoCluster + "/?retryWrites=true&w=majority"

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	collection = client.Database("upvote").Collection("users")
	fmt.Println("Collection is ready")
}

func GetAllUsers() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	var users []primitive.M

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		var user bson.M
		err := cur.Decode(&user)

		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	defer cur.Close(context.Background())

	return users
}

// TO-DO = Actually its returning old value, should return the new value
func NewVote(userId string, upvote bool) primitive.M {

	objId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": objId}
	update := bson.M{"$inc": bson.M{"votes": service.GetVoteValue(upvote)}}

	var updatedUser bson.M
	err = collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&updatedUser)

	// TO-DO = Refact to return error user was not found
	if err != nil {

		if err == mongo.ErrNoDocuments {
			return updatedUser
		}
		log.Fatal(err)

	}

	return updatedUser
}
