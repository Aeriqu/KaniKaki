package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/services/srs/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProviderMongodb struct {
	Client *mongo.Client
}

// -- DATABASE META FUNCTIONS -- //

// Connect attempts to connect to the database server configured via the
// MONGODB_USERNAME and MONGODB_PASSWORD environmental variables.
//
// This will also attempt to connect to do a test connection by pinging it,
// but will exit if the connection is not valid.
func (db *ProviderMongodb) Connect() {
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	uri := fmt.Sprintf("mongodb://%s:%s@mongodb-srs-service.kanikaki.svc.cluster.local:27017/", username, password)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, clientErr := mongo.Connect(context.TODO(), opts)
	if clientErr != nil {
		logger.Fatal("error connecting to mongodb", clientErr)
	}

	db.Client = client
	db.testConnection()
}

// testConnection attempts to test the connection with the database by pinging it.
// The app will exit if the connection is not valid.
func (db *ProviderMongodb) testConnection() bool {
	if runError := db.Client.Ping(context.TODO(), readpref.Primary()); runError != nil {
		logger.Fatal("error pinging mongodb", runError)
		return false
	}

	collection := db.getReviewCollection()
	collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.D{
				{Key: "username", Value: 1},
				{Key: "subject", Value: 1},
				{Key: "topic", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	)

	logger.Info("connection to mongodb successful")
	return true
}

// -- DATABASE ADD FUNCTIONS -- //

// -- DATABASE GET FUNCTIONS -- //

// getReviewCollection grabs the review collection from the database.
func (db *ProviderMongodb) getReviewCollection() *mongo.Collection {
	return db.Client.Database("srs").Collection("review")
}

// GetUserReviewData returns all review data for a user.
func (db *ProviderMongodb) GetUserReviewData(identifier string, subject string) (*[]models.ReviewData, error) {
	collection := db.getReviewCollection()
	filter := bson.M{"username": identifier}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		errMsg := fmt.Sprintf("error finding any review data for the given user (%s): %v", identifier, err)
		logger.Error(errMsg, err)
		return nil, status.Error(codes.NotFound, errMsg)
	}
	defer cursor.Close(context.TODO())

	var results []models.ReviewData
	if err := cursor.All(context.TODO(), &results); err != nil {
		errMsg := fmt.Sprintf("error decoding user review data for (%s): %v", identifier, err)
		logger.Error(errMsg, err)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return &results, nil
}

// GetDueReviewData returns all due reviews for a user and subject.
func (db *ProviderMongodb) GetDueReviewData(identifier string, subject string) (*[]models.ReviewData, error) {
	currentTime := time.Now().Unix()
	collection := db.getReviewCollection()
	filter := bson.M{
		"username":    identifier,
		"subject":     subject,
		"review_time": bson.M{"$lte": currentTime},
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		errMsg := fmt.Sprintf("error finding due review data for user (%s): %v", identifier, err)
		logger.Error(errMsg, err)
		return nil, status.Error(codes.Internal, errMsg)
	}
	defer cursor.Close(context.TODO())

	var results []models.ReviewData
	if err := cursor.All(context.TODO(), &results); err != nil {
		errMsg := fmt.Sprintf("error decoding due review data for user (%s): %v", identifier, err)
		logger.Error(errMsg, err)
		return nil, status.Error(codes.Internal, errMsg)
	}

	return &results, nil
}

// -- DATABASE UPDATE FUNCTIONS -- //

func (db *ProviderMongodb) UpdateSrsItem(identifier string, subject string, topic string, level int32, reviewTime int64) (*models.ReviewData, error) {
	reviewDataFilter := bson.M{
		"username": identifier,
		"subject":  subject,
		"topic":    topic,
	}

	update := bson.M{
		"$set": bson.M{
			"srs_level":   level,
			"review_time": reviewTime,
		},
	}

	findOptions := options.FindOneAndUpdate().
		SetReturnDocument(options.After).
		SetUpsert(true)

	var updatedReviewData models.ReviewData
	collection := db.getReviewCollection()
	err := collection.FindOneAndUpdate(context.TODO(), reviewDataFilter, update, findOptions).Decode(&updatedReviewData)
	if err != nil {
		logger.Error(fmt.Sprintf("error updating srs data (%s - %s - %s)", identifier, subject, topic), err)
		return &updatedReviewData, status.Error(codes.Aborted, fmt.Sprintf("error updating srs data for user (%s)", identifier))
	}

	logger.Info((fmt.Sprintf("updated srs data (%s - %s - %s)", identifier, subject, topic)))

	return &updatedReviewData, nil
}
