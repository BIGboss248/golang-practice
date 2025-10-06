package database

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Sample imports

/*
# MongoDatabaseConnection

# Connect to mongo db and return the connection object

# Packages

go.mongodb.org/mongo-driver/v2/mongo: mongo driver
go.mongodb.org/mongo-driver/v2/mongo/options: mongo driver options

# Parameters

user: string - username.
password: string - password.
host: string - host url.
port: string - port number.

# Returns

MongoDB connection object *mongo.Client
*/
func MongoDatabaseConnection(logger zerolog.Logger, username string, password string, host string, port int) (*mongo.Client, error) {
	DatabaseConnectionStartTime := time.Now() // Record start time
	logger.Debug().Str("FunctionName:", "DatabaseConnection").Msg("DatabaseConnection function started")
	defer func() {
		logger.Debug().Str("FunctionName:", "DatabaseConnection").TimeDiff("Duration (ms)", time.Now(), DatabaseConnectionStartTime).Msg("DatabaseConnection function ended.")
	}()
	if username == "" || password == "" || host == "" || port == 0 {
		err := fmt.Errorf("missing required parameters")
		logger.Error().Err(err).Msg("Missing required parameters")
		return nil, err
	}
	// Build MongoDB URI
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/", username, password, host, port)

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to MongoDB")
		return nil, err
	}
	logger.Info().Msg("✅ Successfully connected to MongoDB")

	return client, nil
}
