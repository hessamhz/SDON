package functions

import (
	"fmt"
	"napcore/internal/client"
)

type InfrastructureParams struct {
	BASE_URL           string
	INFRA_LINE         string
	NE_SRC             string
	NE_DST             string
	CONN_NAME          string
	HIERARCHICAL_LEVEL string
}

func CreateInfra(params InfrastructureParams) (string, error) {

	ID1, err := getNeID(params.BASE_URL, params.INFRA_LINE, params.NE_SRC)
	if err != nil {
		// Handle the error, e.g., log it, return a default value, or exit
		fmt.Println("Error:", err)
		// Handle the error case
	} else {
		// Use the ID
		fmt.Println("ID1:", ID1)
	}

	ID2, err := getNeID(params.BASE_URL, params.INFRA_LINE, params.NE_DST)
	if err != nil {
		// Handle the error, e.g., log it, return a default value, or exit
		fmt.Println("Error:", err)
		// Handle the error case
	} else {
		// Use the ID
		fmt.Println("ID2:", ID2)
	}

	postURLStr := params.BASE_URL + "onc/connection"
	createInfraResponse, err := client.POST(postURLStr, ID1, ID2, params.CONN_NAME, "ConnLpOtu", "otu2x", params.HIERARCHICAL_LEVEL)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return createInfraResponse, nil

}

func getNeID(BASE_URL, INFRA_LINE, neName string) (string, error) {
	urlStr := BASE_URL + "onc/ltp?name==" + INFRA_LINE + "&ne.name==" + neName + "&select(id)"

	ID, err := client.GET(urlStr)
	if err != nil {
		fmt.Println("Error getting client ID:", err)
		return "", err // Return an empty string in case of an error
	}
	fmt.Println(ID)

	return ID, nil
}
