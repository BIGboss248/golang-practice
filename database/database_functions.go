package database

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

/*
# BankAccount

bank acount object represents a bank account that user adds it contains the following fields:

# Fields
 1. AccountNumber string - account number.
 2. CardNumber string - card number.
 3. Name string - name of the account.
 4. Balance float64 - current balance of the account.
 5. Currency string - currency of the account.
 6. BaseAmount float64 - base amount of the account.
 7. Time time.Time - time when the account was created.
 8. Description string - description of the account.
 9. Categories []Category - category of the account.
 10. Tags []string - tags associated with the account.
 11. Transactions []Transaction - list of transactions associated with the account.
*/
type BankAccount struct {
	AccountNumber string
	CardNumber    string
	Name          string
	Balance       float64
	Currency      string
	BaseAmount    float64
	Time          time.Time
	Description   string
	Categories    []Category
	Tags          []string
	Transactions  []Transaction
}

/*
# Transaction

Transaction object represents a transaction that user adds it contains the following fields:

# Fields
 1. ID string - transaction id.
 2. Amount float64 - amount of the transaction.
 3. Time time.Time - time when the transaction was made.
 4. Description string - description of the transaction.
*/
type Transaction struct {
	ID          string
	Amount      float64
	Time        time.Time
	Description string
}

/*
# Catagory

Category object represents a category that user adds it contains the following fields:

# Fields

 1. ID string - category id.
 2. Name string - name of the category.
 3. Description string - description of the category.
*/
type Category struct {
	ID          string
	Name        string
	Description string
}

/*
# MongoDatabaseConnection

# Connect to mongo db and return the connection object

# Packages

1. go.mongodb.org/mongo-driver/v2/mongo: mongo driver

2. go.mongodb.org/mongo-driver/v2/mongo/options: mongo driver options

# Parameters

1. user: string - username.

2. password: string - password.

3. host: string - host url.

4. port: string - port number.

# Returns

1. *mongo.Client: connection object *mongo.Client

2. error: error if any error occurs during connection else nil
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
