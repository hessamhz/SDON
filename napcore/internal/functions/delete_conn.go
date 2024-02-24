package functions

import (
	"fmt"
	"napcore/internal/client"
)

func DeleteConn(BASE_URL string, conn_name string, deleteInfrastructure bool) (bool, error) {
	var deleteResponse bool
	var err error

	if deleteInfrastructure {
		// First Delete Services
		deleteResponse, err = deleteService(BASE_URL, conn_name)
		if err != nil {
			return deleteResponse, err
		}
		fmt.Println(deleteResponse)

		// Delete Infrastructure as well
		infraResponse, infraErr := deleteServiceAndInfrastructure(BASE_URL, conn_name)
		if infraErr != nil {
			return infraResponse, infraErr
		}
		fmt.Println(infraResponse)

	} else {
		// If deleteInfrastructure is false, just delete services
		deleteResponse, err = deleteService(BASE_URL, conn_name)
		if err != nil {
			return deleteResponse, err
		}
	}

	fmt.Println("DeleteResponseHere", deleteResponse)
	return deleteResponse, nil
}

func deleteService(BaseUrl string, conn_name string) (bool, error) {

	urlStrConn := BaseUrl + "onc/connection?name==" + conn_name + "&select(id)"
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
