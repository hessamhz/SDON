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
		BaseUrl:           BASE_URL,
		InfraLine:         "OTU2x-1-1-1",
		NeSrc:             "team1-NE-1",
		NeDst:             "team1-NE-2",
		ConnName:          "FatihConnection",
		HierarchicalLevel: "infrastructure",
	}

	//Only  one 10GB or 	up to 10 1GB Service is allowed

	ServiceParams := functions.ServiceParams{
		BaseUrl:   BASE_URL,
		NeSrc:     "team1-NE-1",
		NeDst:     "team1-NE-2",
		ServRate:  "10Gb",
		NbService: 1,
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
