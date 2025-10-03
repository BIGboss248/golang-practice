package main

// import packages
import (
	"time"

	clog "github.com/BIGboss248/golang-practice/utils/clog"
	"github.com/rs/zerolog"
)

// The function that will be executed
func main() {
	logger, err := clog.SetupLogger("app.log", zerolog.InfoLevel)
	startTime := time.Now() // Record start time
	if err != nil {
		panic(err)
	}
	logger.Debug().Str("FunctionName:", "main").Msg("Main function started")
	defer func() {
		logger.Debug().Str("FunctionName:", "main").TimeDiff("Duration (ms)", time.Now(), startTime).Msg("Main function ended.")
	}()

}
