package main

// import packages
import (
	"os"
	"strconv"
	"time"

	"github.com/BIGboss248/golang-practice/database"
	clog "github.com/BIGboss248/golang-practice/utils/clog"
	"github.com/rs/zerolog"
)

// The function that will be executed
func main() {
	//! Be sure to check for empty string if the environment variable is not set
	var LogLevel string = os.Getenv("Log_Level") // Read environment variable
	zerologLevel := zerolog.InfoLevel
	if LogLevel == "" {
		zerologLevel = zerolog.InfoLevel
	}
	if LogLevel == "Debug" {
		zerologLevel = zerolog.DebugLevel
	}
	logger, err := clog.SetupLogger("app.log", zerologLevel)
	startTime := time.Now() // Record start time
	if err != nil {
		panic(err)
	}
	logger.Debug().Str("FunctionName:", "main").Msg("Main function started")
	defer func() {
		logger.Debug().Str("FunctionName:", "main").TimeDiff("Duration (ms)", time.Now(), startTime).Msg("Main function ended.")
	}()
	portInt, err := strconv.Atoi(os.Getenv("MONGO_PORT"))
	if err != nil {
		logger.Fatal().Err(err).Msg("Invlaid port number")
		return
	}
	database.MongoDatabaseConnection(
		logger,
		os.Getenv("MONGO_USER"),
		os.Getenv("MONGO_PASS"),
		os.Getenv("MONGO_HOST"),
		portInt)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to MongoDB")
	}

}
