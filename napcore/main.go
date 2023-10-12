package main

import (
	"fmt"
	"napcore/env"
	"napcore/internal/abstractions"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	env := env.SetEnv()
	BASE_URL := env.BaseURL

	fmt.Println("Interfaces:")
	params := abstractions.InfrastructureParams{
		BASE_URL:          BASE_URL,
		Infra_Line:        "OTU2x-1-1-1",
		NE_SRC:            "team1-NE-1",
		NE_DST:            "team1-NE-2",
		ConnName:          "FatihConnection",
		HierarchicalLevel: "infrastructure",
	}
	//Only  one 10GB or 	up to 10 1GB Service is allowed

	service_params := abstractions.ServiceParams{
		BASE_URL:   BASE_URL,
		NE_SRC:     "team1-NE-1",
		NE_DST:     "team1-NE-2",
		SERV_RATE:  "10Gb",
		NB_SERVICE: 1,
	}
	fmt.Println(params)
	abstractions.CreateInfra(params)

	abstractions.CreateLP(service_params)

	//abstractions.DeleteConn(service_params, params)

}
