package functions

import (
	"fmt"
	"napcore/internal/client"
	"regexp"
	"strconv"
)

type ServiceParams struct {
	BaseUrl   string
	NeSrc     string
	NeDst     string
	ServRate  string
	NbService int
}

func CreateLP(Params ServiceParams) string {
	// Get ports for Port_Src and Port_Dst

	if Params.ServRate == "10Gb" {
		// Check if NB_SERVICE is not equal to 1
		if Params.NbService != 1 {
			fmt.Println("Error: Only one 10Gb service is allowed.")
			return "Error: Only one 10Gb service is allowed."
		}
	} else if Params.ServRate == "1Gb" {
		// Check if NB_SERVICE is greater than 10
		if Params.NbService > 10 {
			fmt.Println("Error: NB_SERVICE can be at most 10 for 1Gb service.")
			return "Error: NB_SERVICE can be at most 10 for 1Gb service."
		} else if Params.NbService < 1 {
			fmt.Println("Error: NB_SERVICE must be at least 1 for 1Gb service.")
			return "Error: NB_SERVICE must be at least 1 for 1Gb service."
		}
	} else {
		fmt.Println("Error: Unsupported service rate. Only '1Gb' and '10Gb' are allowed.")
		return "Error: Unsupported service rate. Only '1Gb' and '10Gb' are allowed."
	}

	// If the code reaches this point, it means the conditions are met, and you can proceed with the rest of the code.

	portSrc := getPorts(Params.BaseUrl, Params.NeSrc, Params.ServRate)
	portDst := getPorts(Params.BaseUrl, Params.NeDst, Params.ServRate)
	fmt.Println("portSrc", portSrc)
	fmt.Println("portDst", portDst)

	portCommon := getCommonPorts(portSrc, portDst, Params.NbService)
	fmt.Println("Common Ports", portCommon)

	portIDs := getPortIDs(portCommon, Params.BaseUrl, Params.NeSrc, Params.NeDst)
	fmt.Println("portIDs:", portIDs)

	postResponse, err := postPortIDs(portIDs, Params)
	if err != nil {
		fmt.Println("Error:", err)
		return "Error has occured for postPortIDs function"
	}

	return strconv.Itoa(Params.NbService) + "Service Created with each rate:" + Params.ServRate + postResponse
}

func getPorts(BaseURL, neName string, serviceState string) []string {
	fmt.Println("Ports")

	var urlStr string
	if serviceState == "1Gb" {
		urlStr = BaseURL + "onc/ltp?ltpType==physical&ne.name==" + neName + "&select(id,name)&name==OGBE1-*"

	} else {
		urlStr = BaseURL + "onc/ltp?ltpType==physical&ne.name==" + neName + "&select(id,name)&name==OGBE10-*"
	}

	ports, err := client.GET(urlStr)
	if err != nil {
		fmt.Println("Error getting ports:", err)
		return nil
	}
	fmt.Println("Fetched Ports:", ports)
	// Define a regular expression pattern to match "name" values

	pattern := `"name":"([^"]+)"`

	// Compile the regular expression pattern
	re := regexp.MustCompile(pattern)

	// Find all matches in the input string
	matches := re.FindAllStringSubmatch(ports, -1)

	var names []string

	for _, match := range matches {
		if len(match) >= 2 {
			name := match[1]
			names = append(names, name)
		}
	}

	return names
}

func getCommonPorts(portSrc []string, portDst []string, numberOfService int) []string {
	common := []string{}

	// Use a map to store the count of each element in port_src
	count := make(map[string]int)
	for _, item := range portSrc {
		count[item]++
	}

	// Iterate through port_dst and check if the item exists in port_src
	for _, item := range portDst {
		if count[item] > 0 {
			common = append(common, item)
			if len(common) == numberOfService {
				break
			}
		}
	}

	return common
}

func getPortIDs(commonPorts []string, BaseUrl string, neSrc string, neDst string) [][]string {

	fmt.Println("Get by port")

	var portPairs [][]string

	for _, X := range commonPorts {
		// Build the URL strings for Src and Dst
		urlStrSrc := BaseUrl + "onc/ltp?name==" + X + "&ne.name==" + neSrc + "&select(id)"
		urlStrDst := BaseUrl + "onc/ltp?name==" + X + "&ne.name==" + neDst + "&select(id)"
		portSrcID, err := client.GET(urlStrSrc)
		if err != nil {
			fmt.Println("Error getting client ID:", err)
			return nil // Return an empty string in case of an error
		}
		portDstID, err := client.GET(urlStrDst)
		if err != nil {
			fmt.Println("Error getting client ID:", err)
			return nil // Return an empty string in case of an error
		}

		// Create a pair and append it to the array
		portPair := []string{portSrcID, portDstID}
		portPairs = append(portPairs, portPair)

	}

	return portPairs

}

func postPortIDs(portIDs [][]string, Params ServiceParams) (string, error) {

	var postResponse string
	postUrlStr := Params.BaseUrl + "onc/connection"
	for i := 0; i < Params.NbService; i++ {
		serviceName := "Service_" + strconv.Itoa(i+1)
		serviceInfo := "posting to " + postUrlStr + " from source port ID: " + portIDs[i][0] + " to destination port ID: " + portIDs[i][1] + " with service name: " + serviceName
		fmt.Println(serviceInfo)
		postResponse, err := client.POST(postUrlStr, portIDs[i][0], portIDs[i][1], serviceName, "ConnLpEthCbr", Params.ServRate, "service")
		if err != nil {
			fmt.Println("Error", err)
			return "", err
		} else {
			return postResponse, err
		}

	}
	return postResponse, nil
	//client.POST(postUrlStr, Port_IDs[0][0], Port_IDs[0][1], "TestSecSrvc", "ConnLpEthCbr", "1Gb", "service")
	//client.POST(postUrlStr, Port_IDs[0][0], Port_IDs[0][1], "TestSecSrvc", "ConnLpEthCbr", "10Gb", "service")
}
