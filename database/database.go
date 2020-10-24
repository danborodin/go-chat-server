package database

import (
	"context"
	"fmt"
	"log"

	"github.com/danborodin/go-chat-server/config"
	"github.com/danborodin/go-chat-server/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// DB is the database pointer variable
	//DB *mongo.Database

	// DbName - database name
	DbName string = config.Config("DATABASE_NAME")
	// ConnectionString - string for connecting to database
	ConnectionString string = config.Config("DATABASE_CONNECTION_STRING")
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

// AddUserToDatabase add a user to database
func AddUser(user models.User) error {

	client := Connect(ConnectionString)
	userCollection := client.Database(fmt.Sprintf("%s", DbName)).Collection("users")

	//res = insert result
	_, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err)
	}

	defer client.Disconnect(context.TODO())

	return nil
}

func GetUserByID(id string) {

}
