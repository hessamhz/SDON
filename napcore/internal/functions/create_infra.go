package functions

import (
	"fmt"
	"napcore/internal/client"
)

type InfrastructureParams struct {
	BaseUrl           string
	InfraLine         string
	NeSrc             string
	NeDst             string
	ConnName          string
	HierarchicalLevel string
}

func CreateInfra(params InfrastructureParams) (string, error) {

	ID1, err := getNeID(params.BaseUrl, params.InfraLine, params.NeSrc)
	if err != nil {
		// Handle the error, e.g., log it, return a default value, or exit
		fmt.Println("Error:", err)
		// Handle the error case
	} else {
		// Use the ID
		fmt.Println("ID1:", ID1)
	}

	ID2, err := getNeID(params.BaseUrl, params.InfraLine, params.NeDst)
	if err != nil {
		// Handle the error, e.g., log it, return a default value, or exit
		fmt.Println("Error:", err)
		// Handle the error case
	} else {
		// Use the ID
		fmt.Println("ID2:", ID2)
	}

	postURLStr := params.BaseUrl + "onc/connection"
	createInfraResponse, err := client.POST(postURLStr, ID1, ID2, params.ConnName, "ConnLpOtu", "otu2x", params.HierarchicalLevel)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return createInfraResponse, nil

}

func getNeID(BaseUrl, InfraLine, neName string) (string, error) {
	urlStr := BaseUrl + "onc/ltp?name==" + InfraLine + "&ne.name==" + neName + "&select(id)"

	ID, err := client.GET(urlStr)
	if err != nil {
		fmt.Println("Error getting client ID:", err)
		return "", err // Return an empty string in case of an error
	}
	fmt.Println(ID)

	return ID, nil
}
