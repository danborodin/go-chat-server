package database

import (
	"context"
	"fmt"
	"log"

	"github.com/danborodin/go-chat-server/config"
	"github.com/danborodin/go-chat-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// AddRoom add new room to database
func AddRoom(room models.Room) error {
	client := Connect(ConnectionString)
	roomsCollection := client.Database(fmt.Sprintf("%s", DbName)).Collection("rooms")

	res, err := roomsCollection.InsertOne(context.TODO(), room)
	if err != nil {
		log.Println("Error adding a room, ", err)
		return err
	}

	log.Println(fmt.Sprintf("Room with id %v added", res.InsertedID))

	defer client.Disconnect(context.TODO())

	return err
}

func GetRoomById(id string) (models.Room, error) {
	client := Connect(ConnectionString)
	roomsCollection := client.Database(fmt.Sprintf("%s", DbName)).Collection("rooms")
	_id, err := primitive.ObjectIDFromHex(id)
	log.Println("huynea: ", _id)
	result := roomsCollection.FindOne(context.Background(), bson.M{"_id": _id})

	room := models.Room{}
	result.Decode(room)
	log.Println(room)

	return room, err
}

// working...

// GetRooms ...
func GetRooms() ([]models.Room, error) {
	client := Connect(ConnectionString)
	roomsCollection := client.Database(fmt.Sprintf("%s", DbName)).Collection("rooms")

	var result []models.Room
	cursor, err := roomsCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		log.Println(err)
		return result, err
	}
	defer client.Disconnect(context.TODO())

	return result, err
}
