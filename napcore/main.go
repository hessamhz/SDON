package main

import (
	"fmt"
	"napcore/env"
	"napcore/internal/functions"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	env := env.SetEnv()
	BASE_URL := env.BaseURL

	Params := functions.InfrastructureParams{
		BASE_URL:           BASE_URL,
		INFRA_LINE:         "OTU2x-1-1-1",
		NE_SRC:             "team1-NE-1",
		NE_DST:             "team1-NE-2",
		CONN_NAME:          "FatihConnection",
		HIERARCHICAL_LEVEL: "infrastructure",
	}

	//Only  one 10GB or 	up to 10 1GB Service is allowed

	ServiceParams := functions.ServiceParams{
		BASE_URL:   BASE_URL,
		NE_SRC:     "team1-NE-1",
		NE_DST:     "team1-NE-2",
		SERV_RATE:  "10Gb",
		NB_SERVICE: 1,
	}

	DeleteInfrastructureAsWell := true
	/*
			CreateInfraResponse, err := functions.CreateInfra(Params)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				fmt.Println("CreateInfraResponse", CreateInfraResponse)
			}

			fmt.Println("Infrastructure Created now its turn for CreateLPResponse ")


		CreateLPResponse := functions.CreateLP(ServiceParams)
		fmt.Println("CreateLPResponse", CreateLPResponse)
		fmt.Println("Lightpath Created now its turn for Delete Connections ")
	*/
	DeleteResponse, err := functions.DeleteConn(ServiceParams, Params, DeleteInfrastructureAsWell)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("DeleteResponse", DeleteResponse)
	}

}
