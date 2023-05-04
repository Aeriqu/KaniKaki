package database

import (
	"context"
	"fmt"
	"os"

	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/common/models"
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
// The host will be mongodb-auth-service.kanikaki.svc.cluster.local:27017
//
// This will also attempt to connect to do a test connection by pinging it,
// but will exit if the connection is not valid.
func (db *ProviderMongodb) Connect() {
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	uri := fmt.Sprintf("mongodb://%s:%s@mongodb-kanji-service.kanikaki.svc.cluster.local:27017/", username, password)
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

	collection := db.getKanjiCollection()
	collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "character", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	logger.Info("connection to mongodb successful")
	return true
}

// -- DATABASE ADD FUNCTIONS -- //

// AddMultipleKanji adds a list of kanji to the database.
func (db *ProviderMongodb) AddMultipleKanji(kanjiList []models.Kanji) (string, error) {
	collection := db.getKanjiCollection()

	insertInterface := make([]interface{}, len(kanjiList))
	insertOptions := options.InsertMany().SetOrdered(false)

	for index := range kanjiList {
		insertInterface[index] = kanjiList[index]
	}

	insert, err := collection.InsertMany(context.TODO(), insertInterface, insertOptions)
	if err != nil {
		errMsg := "error inserting kanji into the collection (it may already exist)"
		logger.Error(errMsg, err)
		return "", status.Error(codes.Aborted, errMsg)
	}

	insertResult, _ := bson.MarshalExtJSON(insert, false, false)

	logger.Info((fmt.Sprintf("added kanji (%d)", len(kanjiList))))

	return string(insertResult), nil
}

// -- DATABASE GET FUNCTIONS -- //

// getKanjiCollection grabs the user collection from the auth database.
func (db *ProviderMongodb) getKanjiCollection() *mongo.Collection {
	return db.Client.Database("auth").Collection("kanji")
}

// GetKanjiByCharacter grabs a kanji with the a given identifier (character).
func (db *ProviderMongodb) GetKanjiByCharacter(character string) (*models.Kanji, error) {
	results := &[]models.Kanji{}
	collection := db.getKanjiCollection()

	collectionCursor, err := collection.Find(context.TODO(),
		bson.D{
			{Key: "character", Value: character},
		},
	)

	if err != nil {
		logger.Error(fmt.Sprintf("error finding kanji (%s)", character), err)
		return nil, err
	}

	if err = collectionCursor.All(context.TODO(), results); err != nil {
		errMsg := fmt.Sprintf("error selecting kanji (%s)", character)
		logger.Error(errMsg, err)
		return nil, status.Error(codes.Unknown, errMsg)
	}

	if len(*results) == 0 {
		return nil, status.Error(codes.NotFound, "no kanji found")
	} else if len(*results) > 1 {
		errMsg := fmt.Sprintf("error found too many kanji (%d)", len(*results))
		logger.Error(errMsg, err)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	return &(*results)[0], nil
}

// GetKanjiByRange grabs a list a kanji by a given range
func (db *ProviderMongodb) GetKanjiByRange(lowerLimit int, upperLimit int) (*[]models.Kanji, error) {
	results := &[]models.Kanji{}
	collection := db.getKanjiCollection()

	findFilter := bson.M{
		"wanikani_level": bson.M{
			"$gte": lowerLimit,
			"$lte": upperLimit,
		},
	}

	collectionCursor, err := collection.Find(context.TODO(),
		findFilter,
	)
	if err != nil {
		errMsg := fmt.Sprintf("error finding kanji in the range (%d - %d)", lowerLimit, upperLimit)
		logger.Error(errMsg, err)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	if err = collectionCursor.All(context.TODO(), results); err != nil {
		errMsg := fmt.Sprintf("error selecting kanji in the range (%d - %d)", lowerLimit, upperLimit)
		logger.Error(errMsg, err)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	return results, nil
}
