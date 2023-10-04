package main

import (
	"napcore/env"
	"napcore/internal/client"

	"github.com/joho/godotenv"

)

func main() {
	godotenv.Load()

	env := env.SetEnv()

	urlStr := env.BaseURL + "onc/ltp?name==OTU2x-1-1-1&ne.name==team1-NE-2&select(id)"
	postUrlStr := env.BaseURL + "onc/connection"
	client.GetInterface(urlStr)
	client.CreateLP(postUrlStr)
}
