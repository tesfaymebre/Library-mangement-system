package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	userCollection   *mongo.Collection
	taskCollection   *mongo.Collection
	signupCollection *mongo.Collection
)

// InitializeMongoDB establishes a connection to MongoDB
func InitializeMongoDB() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// get mongodb uri from .env file
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI not set in .env file")
	}

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// Get a handle for your collection
	userCollection = client.Database("taskdb").Collection("users")
	taskCollection = client.Database("taskdb").Collection("tasks")
	signupCollection = client.Database("taskdb").Collection("signup")
}
