package functions

import (
	"fmt"
	"napcore/internal/client"
)

func DeleteConn(ServParams ServiceParams, InfParams InfrastructureParams, deleteInfrastructure bool) (bool, error) {
	var deleteResponse bool
	var err error

	if deleteInfrastructure {
		// First Delete Services
		deleteResponse, err = deleteService(ServParams.BaseUrl, ServParams.NbService)
		if err != nil {
			return deleteResponse, err
		}
		fmt.Println(deleteResponse)

		// Delete Infrastructure as well
		infraResponse, infraErr := deleteServiceAndInfrastructure(ServParams.BaseUrl, InfParams.ConnName)
		if infraErr != nil {
			return infraResponse, infraErr
		}
		fmt.Println(infraResponse)

	} else {
		// If deleteInfrastructure is false, just delete services
		deleteResponse, err = deleteService(ServParams.BaseUrl, ServParams.NbService)
		if err != nil {
			return deleteResponse, err
		}
	}

	fmt.Println("DeleteResponseHere", deleteResponse)
	return deleteResponse, nil
}

func deleteService(BaseUrl string, NbService int) (bool, error) {

	for i := 1; i <= NbService; i++ {
		serviceNum := fmt.Sprintf("%d", i)

		urlStrConn := BaseUrl + "onc/connection?name==Service_" + serviceNum + "&select(id)"
		connID, err := client.GET(urlStrConn)
		if err != nil {
			fmt.Println("Error getting client ID:", err)
			return false, err
		}

		urlStrConnSerDel := BaseUrl + "onc/connection/" + connID
		fmt.Println("CONNID", urlStrConnSerDel)
		deleteResponse, err := client.DELETE(urlStrConnSerDel)
		if err != nil {
			fmt.Println("Deleting Error:", err)
			return deleteResponse, err
		}

	}

	return true, nil

}

func deleteServiceAndInfrastructure(BaseUrl string, connName string) (bool, error) {

	urlStrConnInfra := BaseUrl + "onc/connection?name==" + connName + "&select(id)"
	infraConnID, err := client.GET(urlStrConnInfra)
	if err != nil {
		fmt.Println("Error getting client ID:", err)
		return false, err
	}

	urlStrConnInfraDel := BaseUrl + "onc/connection/" + infraConnID
	fmt.Println("CONNID", urlStrConnInfraDel)

	deleteResponse, err := client.DELETE(urlStrConnInfraDel)
	if err != nil {
		fmt.Println("Deleting Error:", err)
		return false, err
	} else {
		fmt.Print(`Delete Response:`, deleteResponse)

	}

	return deleteResponse, nil

}
