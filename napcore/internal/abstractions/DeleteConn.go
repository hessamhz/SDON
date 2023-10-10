package abstractions

import (
	"fmt"
	"napcore/internal/client"
)

func DeleteConn(Serv_Params ServiceParams, Inf_Params InfrastructureParams) {

	getConnID(Serv_Params.BASE_URL, Inf_Params.ConnName, Serv_Params.NB_SERVICE)

}

func getConnID(BASE_URL string, ConnName string, NB_SERVICE int) {

	for i := 1; i <= NB_SERVICE; i++ {
		ServiceNum := fmt.Sprintf("%d", i)

		urlStrConn := BASE_URL + "onc/connection?name==Service_" + ServiceNum + "&select(id)"
		ConnID, err := client.GET(urlStrConn)
		if err != nil {
			fmt.Println("Error getting client ID:", err)
			return // Return an empty string in case of an error
		}

		urlStrConnSerDel := BASE_URL + "onc/connection/" + ConnID
		fmt.Println("CONNID", urlStrConnSerDel)
		client.DELETE(urlStrConnSerDel)

	}

	urlStrConnInfra := BASE_URL + "onc/connection?name==" + ConnName + "&select(id)"
	InfraConnID, err := client.GET(urlStrConnInfra)
	if err != nil {
		fmt.Println("Error getting client ID:", err)
		return // Return an empty string in case of an error
	}
	urlStrConnInfraDel := BASE_URL + "onc/connection/" + InfraConnID
	fmt.Println("CONNID", urlStrConnInfraDel)
	client.DELETE(urlStrConnInfraDel)

}
