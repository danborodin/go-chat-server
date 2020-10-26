package database

import (
	"context"
	"fmt"
	"log"

	"github.com/danborodin/go-chat-server/config"
	"github.com/danborodin/go-chat-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// DB is the database pointer variable
	//DB *mongo.Database

	// DbName - database name
	DbName string = config.GetEnvVar("DATABASE_NAME")
	// ConnectionString - string for connecting to database
	ConnectionString string = config.GetEnvVar("DATABASE_CONNECTION_STRING")
	// Context is the context variable
	//Context context.Context, Context2 context.CancelFunc() = context.WithTimeout(context.Background(), 10*time.Second)
)

// Connect to database
func Connect(connString string) *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		log.Println(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Println(err)
	}

	return client
}

// AddUser add a user to database
func AddUser(user models.User) error {

	client := Connect(ConnectionString)
	userCollection := client.Database(fmt.Sprintf("%s", DbName)).Collection("users")

	//res = insert result
	res, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(fmt.Sprintf("User with id %v added", res.InsertedID))

	defer client.Disconnect(context.TODO())

	return err
}

// GetUserByUsername ...
func GetUserByUsername(username string) (models.User, error) {
	client := Connect(ConnectionString)
	userCollection := client.Database(fmt.Sprintf("%s", DbName)).Collection("users")

	var result models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&result)
	if err != nil {
		return result, err
	}
	defer client.Disconnect(context.TODO())

	return result, err
}

// working...
// GetChannels ...
func GetChannels() ([]models.Channel, error) {
	client := Connect(ConnectionString)
	channelCollection := client.Database(fmt.Sprintf("%s", DbName)).Collection("channels")

	var result []models.Channel
	err := channelCollection.FindOne(context.TODO(), bson.M{"name": "channelu_pulii_mele"}).Decode(&result)
	if err != nil {
		log.Println(err)
		return result, err
	}
	// err = cursor.All(context.TODO(), &result)
	// if err != nil {
	// 	log.Println(err)
	// 	return result, err
	// }
	defer client.Disconnect(context.TODO())

	return result, err
}
