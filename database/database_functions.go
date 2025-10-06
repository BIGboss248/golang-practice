package database

import (
	"context"
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
	// Build MongoDB URI
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/", username, password, host, port)

	// Define connection options
	clientOptions := options.Client().ApplyURI(uri)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("could not ping MongoDB: %w", err)
	}

	logger.Info().Msg("✅ Connected to MongoDB: " + username + "@" + host + ":" + fmt.Sprint(port))
	return client, nil

}
