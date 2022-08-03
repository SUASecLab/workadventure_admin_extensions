package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	Username string
	Password string
	Hostname string
	Dbname   string
}

func (database MongoDatabase) QueryUserInformation(queryType QueryType, uuid string) (bool, error) {
	if queryType != UserExists {
		return false, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+database.Username+":"+
		database.Password+"@"+database.Hostname+":27017/"+database.Dbname))
	defer client.Disconnect(ctx)

	if err != nil {
		return false, err
	}

	collection := client.Database(database.Dbname).Collection("users")

	var result bson.M

	err = collection.FindOne(ctx, bson.D{{"uuid", uuid}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
	} else {
		// if we find a document matching the uuid, the user exists
		return true, nil
	}

	return false, nil
}
