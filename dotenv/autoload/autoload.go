package autoload

import (
	"os"

	"github.com/joho/godotenv"

	dotenvex "github.com/chouandy/go-sdk/dotenv"
)

func init() {
	// Load for go-gin
	ginMode := os.Getenv("GIN_MODE")
	if len(ginMode) > 0 {
		godotenv.Load(".env." + ginMode)
	}

	// Load by stage
	dotenvex.LoadByStage()

	// Load .env
	godotenv.Load()
}
