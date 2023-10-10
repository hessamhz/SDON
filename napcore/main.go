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

	service_params := abstractions.ServiceParams{
		BASE_URL:   BASE_URL,
		NE_SRC:     "team1-NE-1",
		NE_DST:     "team1-NE-2",
		NB_SERVICE: 3,
	}

	//abstractions.CreateInfra(params)

	//abstractions.CreateLP(service_params)

	abstractions.DeleteConn(service_params, params)
	//client.POST(postUrlStr, "94", "152", "TestSecSrvc", "ConnLpEthCbr", "1Gb", "service") // id: 85

	//fmt.Println("Get Service")
	//urlStr = env.BaseURL + "onc/connection?name==FatihConnection&select(id)"
	//client.GET(urlStr)

	//fmt.Println("Deleting")
	//urlStr = env.BaseURL + "onc/connection/78"
	//client.DELETE(urlStr)

}
