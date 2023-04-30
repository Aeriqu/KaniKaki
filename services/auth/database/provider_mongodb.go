package database

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/services/auth/models"
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
	uri := fmt.Sprintf("mongodb://%s:%s@mongodb-auth-service.kanikaki.svc.cluster.local:27017/", username, password)
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

	collection := db.getUserCollection()
	collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	logger.Info("connection to mongodb successful")
	return true
}

// -- DATABASE ADD FUNCTIONS -- //

// AddUser creates a new user in the users collection of the auth database.
// Does a check to see if the user exists before creating it, to enforce
// unique users.
//
// This will return an error if the user already exists.
func (db *ProviderMongodb) AddUser(username string, passwordHash string) (string, error) {
	// Create the user
	collection := db.getUserCollection()
	user := models.User{
		Username:   username,
		Password:   passwordHash,
		AuthTokens: []models.Token{},
	}

	insert, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		errMsg := fmt.Sprintf("error inserting a new user (%s) into the collection (it may already exist)", username)
		logger.Error(errMsg, err)
		return "", errors.New(errMsg)
	}

	insertResult, _ := bson.MarshalExtJSON(insert, false, false)

	logger.Info((fmt.Sprintf("added user (%s): %s", username, insertResult)))

	return string(insertResult), nil
}

// AddTokenToUser adds a given token to a user with a given identifier (username)
func (db *ProviderMongodb) AddTokenToUser(username string, newToken string, newTokenExpiration int64) (string, error) {
	// Add the token to the user
	collection := db.getUserCollection()

	userFilter := bson.D{{Key: "username", Value: username}}
	updateOperation := bson.D{{
		Key: "$push",
		Value: bson.D{{
			Key: "auth_tokens",
			Value: models.Token{
				Jwt:        newToken,
				Expiration: newTokenExpiration,
			},
		}},
	}}

	update, err := collection.UpdateOne(
		context.TODO(),
		userFilter,
		updateOperation,
	)

	if err != nil {
		logger.Error(fmt.Sprintf("error inserting token (%s) for user (%s)", newToken, username), err)
		return "", err
	}

	updateResult, _ := bson.MarshalExtJSON(update, false, false)

	logger.Info((fmt.Sprintf("added new token to user (%s): %s", username, updateResult)))

	return string(updateResult), nil
}

// -- DATABASE GET FUNCTIONS -- //

// getUserCollection grabs the user collection from the auth database.
func (db *ProviderMongodb) getUserCollection() *mongo.Collection {
	return db.Client.Database("auth").Collection("users")
}

// GetUserByIdentifier grabs a user with the a given identifier (username).
func (db *ProviderMongodb) GetUserByIdentifier(username string) (*models.User, error) {
	results := &[]models.User{}
	collection := db.getUserCollection()

	collectionCursor, err := collection.Find(context.TODO(),
		bson.D{
			{Key: "username", Value: username},
		},
	)

	if err != nil {
		logger.Error(fmt.Sprintf("error finding user (%s)", username), err)
		return nil, err
	}

	if err = collectionCursor.All(context.TODO(), results); err != nil {
		errMsg := fmt.Sprintf("error selecting user (%s)", username)
		logger.Error(errMsg, err)
		return nil, status.Error(codes.Unknown, errMsg)
	}

	if len(*results) == 1 {
		return &(*results)[0], nil
	} else if len(*results) > 1 {
		errMsg := fmt.Sprintf("error found too many accounts (%d)", len(*results))
		logger.Error(errMsg, err)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	return nil, status.Error(codes.NotFound, "no accounts found")
}

// GetUserByIdentifier grabs all users with the a given token.
func (db *ProviderMongodb) GetUserByToken(token string) (*models.User, error) {
	results := &[]models.User{}
	collection := db.getUserCollection()

	collectionCursor, err := collection.Find(context.TODO(),
		bson.D{
			{Key: "auth_tokens.jwt", Value: token},
		},
	)

	if err != nil {
		logger.Error("error finding user with token: "+token, err)
		return nil, err
	}

	if err = collectionCursor.All(context.TODO(), results); err != nil {
		errMsg := fmt.Sprintf("error selecting user with token (%s)", token)
		logger.Error(errMsg, err)
		return nil, status.Error(codes.Unknown, errMsg)
	}

	if len(*results) == 1 {
		return &(*results)[0], nil
	} else if len(*results) > 1 {
		errMsg := fmt.Sprintf("error found too many accounts (%d)", len(*results))
		logger.Error(errMsg, err)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	return nil, status.Error(codes.NotFound, "no accounts found")
}

// -- DATABASE UPDATE FUNCTIONS -- //

// UpdateUserIdentifier updates a given user's identifier (username)
// to a new provided identifier (username).
func (db *ProviderMongodb) UpdateUserIdentifier(username string, newUsername string) (string, error) {
	// Add the token to the user
	collection := db.getUserCollection()
	update, err := collection.UpdateOne(
		context.TODO(),
		bson.D{{Key: "username", Value: username}},
		bson.D{{Key: "$set", Value: bson.D{{Key: "username", Value: newUsername}}}},
	)

	if err != nil {
		errMsg := fmt.Sprintf("error updating username (%s) for user (%s)", newUsername, username)
		logger.Error(errMsg, err)
		return "", status.Error(codes.Aborted, errMsg)
	}

	updateResult, _ := bson.MarshalExtJSON(update, false, false)

	logger.Info((fmt.Sprintf("updated user's (%s) username (%s): %s", username, newUsername, updateResult)))

	return string(updateResult), nil
}

// UpdateUserCredential updates a given user's credential (password)
func (db *ProviderMongodb) UpdateUserCredential(username string, newCredential string) (string, error) {
	// Add the token to the user
	collection := db.getUserCollection()
	updateResult, err := collection.UpdateOne(
		context.TODO(),
		bson.D{{Key: "username", Value: username}},
		bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: newCredential}}}},
	)

	if err != nil {
		logger.Error(fmt.Sprintf("error updating credential (%s) for user (%s)", newCredential, username), err)
		return "", status.Error(codes.Aborted, fmt.Sprintf("error updating credential for user (%s)", username))
	}

	response, _ := bson.MarshalExtJSON(updateResult, false, false)

	logger.Info((fmt.Sprintf("updated user credential (%s)", username)))

	return string(response), nil
}

// UpdateSpecificUserToken replaces an oldToken with a newToken on a given username
func (db *ProviderMongodb) UpdateSpecificUserToken(username string, oldToken string, newToken string, newTokenExpiration int64) (*models.User, error) {
	userFilter := bson.D{{Key: "username", Value: username}}

	tokenFilter := []interface{}{bson.M{
		"elem.jwt": oldToken,
	}}

	findOptions := options.FindOneAndUpdate().
		SetArrayFilters(options.ArrayFilters{
			Filters: tokenFilter,
		}).
		SetReturnDocument(options.After)

	updateOperator := bson.D{{
		Key: "$set",
		Value: bson.D{
			{
				Key:   "auth_tokens.$[elem].jwt",
				Value: newToken,
			},
			{
				Key:   "auth_tokens.$[elem].expiration",
				Value: newTokenExpiration,
			},
		},
	}}

	var updatedUser models.User
	collection := db.getUserCollection()
	err := collection.FindOneAndUpdate(context.TODO(), userFilter, updateOperator, findOptions).Decode(&updatedUser)
	if err != nil {
		logger.Error(fmt.Sprintf("error replacing old token with new token (%s) for user (%s)", newToken, username), err)
		return &updatedUser, status.Error(codes.Aborted, fmt.Sprintf("error updating token for user (%s)", username))
	}

	logger.Info((fmt.Sprintf("updated token for user (%s)", username)))

	return &updatedUser, nil
}

// -- DATABASE REMOVE FUNCTIONS -- //

// Removes expired tokens from a given user
func (db *ProviderMongodb) RemoveUserExpiredTokens(username string) (*models.User, error) {
	currentTime := time.Now().Unix()
	userFilter := bson.D{{Key: "username", Value: username}}

	updateOperator := bson.D{{
		Key: "$pull",
		Value: bson.D{{
			Key: "auth_tokens",
			Value: bson.D{{
				Key: "expiration",
				Value: bson.D{{
					Key:   "$lte",
					Value: currentTime,
				}},
			}},
		}},
	}}

	var updatedUser models.User
	collection := db.getUserCollection()
	err := collection.FindOneAndUpdate(context.TODO(), userFilter, updateOperator).Decode(&updatedUser)
	if err != nil {
		logger.Error(fmt.Sprintf("error removing expired tokens for user (%s)", username), err)
		return &updatedUser, status.Error(codes.Aborted, "error removing expired tokens")
	}

	logger.Info((fmt.Sprintf("removed expired tokens from user (%s)", username)))

	return &updatedUser, nil
}

// Removes a tokens from a given user
func (db *ProviderMongodb) RemoveSpecificUserToken(username string, token string) (*models.User, error) {
	userFilter := bson.D{{Key: "username", Value: username}}

	updateOperator := bson.D{{
		Key: "$pull",
		Value: bson.D{{
			Key: "auth_tokens",
			Value: bson.D{{
				Key: "jwt",
				Value: bson.D{{
					Key:   "$eq",
					Value: token,
				}},
			}},
		}},
	}}

	var updatedUser models.User
	collection := db.getUserCollection()
	err := collection.FindOneAndUpdate(context.TODO(), userFilter, updateOperator).Decode(&updatedUser)
	if err != nil {
		logger.Error(fmt.Sprintf("error removing expired tokens for user (%s)", username), err)
		return &updatedUser, status.Error(codes.Aborted, "error removing expired tokens")
	}

	logger.Info((fmt.Sprintf("removed expired tokens from user (%s)", username)))

	return &updatedUser, nil
}