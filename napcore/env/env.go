package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type NapcoreEnv struct {
	BaseURL string
}

func SetEnv() NapcoreEnv {
	godotenv.Load()

	var errs []string

	BaseURL := os.Getenv("BASE_URL")
	if BaseURL == "" {
		errs = append(errs, "BASE_URL Not Found")
	}

	if len(errs) > 0 {
		log.Fatal("Variable(s) have issues: ", errs)
	}

	return NapcoreEnv{
		BaseURL: BaseURL,
	}
}
