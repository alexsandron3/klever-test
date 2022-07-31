package model

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var collection *mongo.Collection

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
func UpdateVote(voteValue int16, objId primitive.ObjectID) (primitive.M, error) {

	filter := bson.M{"_id": objId}
	update := bson.M{"$inc": bson.M{"votes": voteValue}}

	var updatedUser bson.M
	err := collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&updatedUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "User not found")
		}
		return nil, status.Errorf(codes.Internal, "An internal error ocurred")
	}

	return updatedUser, nil
}
