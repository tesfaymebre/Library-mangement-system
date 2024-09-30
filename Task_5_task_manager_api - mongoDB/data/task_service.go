package data

import (
	"context"
	"errors"
	"log"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient *mongo.Client
	collection  *mongo.Collection
)

// InitializeMongoDB establishes a connection to MongoDB
func InitializeMongoDB() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get a handle for your collection
	collection = client.Database("task_manager").Collection("tasks")
}

// getTasks responds with the list of all tasks as JSON.
func GetTasks() []models.Task {
	var tasks []models.Task
	cur, err := collection.Find(context.TODO(), bson.D{{}}) // bson.D{} specifies 'all documents'

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var task models.Task
		err := cur.Decode(&task)

		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, task)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return tasks
}

// getTask responds with the details of a specific task.
func GetTaskById(taskID string) (models.Task, error) {
	var task models.Task
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return models.Task{}, errors.New("invalid task ID")
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&task)

	if err != nil {
		return models.Task{}, errors.New("task not found")
	}

	return task, nil
}

// createTask creates a new task.
func CreateTask(newTask models.Task) models.Task {
	_, err := collection.InsertOne(context.TODO(), newTask)

	if err != nil {
		log.Fatal(err)
	}

	return newTask
}

// updateTask updates a specific task.
func UpdateTask(taskID string, updatedTask models.Task) (models.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return models.Task{}, errors.New("invalid task ID")
	}

	filter := bson.M{"_id": objectID}

	update := bson.M{

		"$set": updatedTask,
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var result models.Task

	err = collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&result)

	if err != nil {

		if err == mongo.ErrNoDocuments {

			return models.Task{}, errors.New("task not found")

		}

		return models.Task{}, err

	}

	return result, nil
}

// deleteTask deletes a specific task by ID.
func DeleteTask(taskID string) error {
	objectID, err := primitive.ObjectIDFromHex(taskID)

	if err != nil {
		return errors.New("invalid task ID")
	}

	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})

	if err != nil {
		log.Fatal(err)
	}

	if deleteResult.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
