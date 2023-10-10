package abstractions

import (
	"fmt"
	"napcore/internal/client"
)

type InfrastructureParams struct {
	BASE_URL          string
	Infra_Line        string
	NE_SRC            string
	NE_DST            string
	ConnName          string
	HierarchicalLevel string
}

func CreateInfra(params InfrastructureParams) {

	ID1 := getNeID(params.BASE_URL, params.Infra_Line, params.NE_SRC)
	ID2 := getNeID(params.BASE_URL, params.Infra_Line, params.NE_DST)

	postURLStr := params.BASE_URL + "onc/connection"
	client.POST(postURLStr, ID1, ID2, params.ConnName, "ConnLpOtu", "otu2x", params.HierarchicalLevel)

}

func getNeID(BASE_URL, Infra_Line, NEName string) string {
	urlStr := BASE_URL + "onc/ltp?name==" + Infra_Line + "&ne.name==" + NEName + "&select(id)"

	ID, err := client.GET(urlStr)
	if err != nil {
		fmt.Println("Error getting client ID:", err)
		return "xxx" // Return an empty string in case of an error
	}
	fmt.Println(ID)

	return ID
}
