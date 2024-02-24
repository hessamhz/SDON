package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type NapcoreEnv struct {
	BaseURL      string
	InfluxURL    string
	InfluxToken  string
	InfluxBucket string
	InfluxOrg    string
}

func SetEnv() NapcoreEnv {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	var missingVars []string

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		missingVars = append(missingVars, "BASE_URL")
	}

	influxURL := os.Getenv("INFLUX_URL")
	if influxURL == "" {
		missingVars = append(missingVars, "INFLUX_URL")
	}

	influxToken := os.Getenv("INFLUX_TOKEN")
	if influxToken == "" {
		missingVars = append(missingVars, "INFLUX_TOKEN")
	}

	influxBucket := os.Getenv("INFLUX_BUCKET")
	if influxBucket == "" {
		missingVars = append(missingVars, "INFLUX_BUCKET")
	}

	influxOrg := os.Getenv("INFLUX_ORG")
	if influxOrg == "" {
		missingVars = append(missingVars, "INFLUX_ORG")
	}

	if len(missingVars) > 0 {
		log.Fatalf("Missing required environment variable(s): %v", missingVars)
	}

	return NapcoreEnv{
		BaseURL:      baseURL,
		InfluxURL:    influxURL,
		InfluxToken:  influxToken,
		InfluxBucket: influxBucket,
		InfluxOrg:    influxOrg,
	}
}
