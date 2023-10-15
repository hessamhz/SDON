package functions

import (
	"fmt"
	"napcore/internal/client"
)

func DeleteConn(Serv_Params ServiceParams, Inf_Params InfrastructureParams, deleteInfrastructure bool) (string, error) {
	var deleteResponse string
	var err error

	if deleteInfrastructure {
		// First Delete Services
		deleteResponse, err = deleteService(Serv_Params.BASE_URL, Inf_Params.CONN_NAME, Serv_Params.NB_SERVICE)
		if err != nil {
			return deleteResponse, err
		}
		fmt.Println(deleteResponse)

		// Delete Infrastructure as well
		infraResponse, infraErr := deleteServiceAndInfrastructure(Serv_Params.BASE_URL, Inf_Params.CONN_NAME)
		if infraErr != nil {
			return infraResponse, infraErr
		}
		fmt.Println(infraResponse)

	} else {
		// If deleteInfrastructure is false, just delete services
		deleteResponse, err = deleteService(Serv_Params.BASE_URL, Inf_Params.CONN_NAME, Serv_Params.NB_SERVICE)
		if err != nil {
			return deleteResponse, err
		}
	}

	fmt.Println("DeleteResponseHere", deleteResponse)
	return deleteResponse, nil
}

func deleteService(BASE_URL string, connName string, NB_SERVICE int) (string, error) {

	var deleteResponse string

	for i := 1; i <= NB_SERVICE; i++ {
		ServiceNum := fmt.Sprintf("%d", i)

		urlStrConn := BASE_URL + "onc/connection?name==Service_" + ServiceNum + "&select(id)"
		connID, err := client.GET(urlStrConn)
		if err != nil {
			fmt.Println("Error getting client ID:", err)
			return "", err
		}

		urlStrConnSerDel := BASE_URL + "onc/connection/" + connID
		fmt.Println("CONNID", urlStrConnSerDel)
		deleteResponse, err := client.DELETE(urlStrConnSerDel)
		if err != nil {
			fmt.Println("Deleting Error:", err)
			return deleteResponse, err
		}

	}

	return deleteResponse, nil

}

func deleteServiceAndInfrastructure(BASE_URL string, connName string) (string, error) {

	urlStrConnInfra := BASE_URL + "onc/connection?name==" + connName + "&select(id)"
	infraConnID, err := client.GET(urlStrConnInfra)
	if err != nil {
		fmt.Println("Error getting client ID:", err)
		return "", err
	}

	urlStrConnInfraDel := BASE_URL + "onc/connection/" + infraConnID
	fmt.Println("CONNID", urlStrConnInfraDel)

	deleteResponse, err := client.DELETE(urlStrConnInfraDel)
	if err != nil {
		fmt.Println("Deleting Error:", err)
		return "", err
	} else {
		fmt.Print(`Delete Response:`, deleteResponse)

	}

	return deleteResponse, nil

}
