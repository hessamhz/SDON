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

func CreateLP(params ServiceParams) string {
	// Get ports for Port_Src and Port_Dst

	if params.ServRate == "10Gb" {
		// Check if NB_SERVICE is not equal to 1
		if params.NbService != 1 {
			fmt.Println("Error: Only one 10Gb service is allowed.")
			return "Error: Only one 10Gb service is allowed."
		}
	} else if params.ServRate == "1Gb" {
		// Check if NB_SERVICE is greater than 10
		if params.NbService > 10 {
			fmt.Println("Error: NB_SERVICE can be at most 10 for 1Gb service.")
			return "Error: NB_SERVICE can be at most 10 for 1Gb service."
		} else if params.NbService < 1 {
			fmt.Println("Error: NB_SERVICE must be at least 1 for 1Gb service.")
			return "Error: NB_SERVICE must be at least 1 for 1Gb service."
		}
	} else {
		fmt.Println("Error: Unsupported service rate. Only '1Gb' and '10Gb' are allowed.")
		return "Error: Unsupported service rate. Only '1Gb' and '10Gb' are allowed."
	}

	// If the code reaches this point, it means the conditions are met, and you can proceed with the rest of the code.

	portSrc := getPorts(params.BaseUrl, params.NeSrc, params.ServRate)
	portDst := getPorts(params.BaseUrl, params.NeDst, params.ServRate)
	fmt.Println("portSrc", portSrc)
	fmt.Println("portDst", portDst)

	portCommon := getCommonPorts(portSrc, portDst, params.NbService)
	fmt.Println("Common Ports", portCommon)

	portIDs := getPortIDs(portCommon, params.BaseUrl, params.NeSrc, params.NeDst)
	fmt.Println("portIDs:", portIDs)

	postResponse, err := postPortIDs(portIDs, params)
	if err != nil {
		fmt.Println("Error:", err)
		return "Error has occured for postPortIDs function"
	}

	return strconv.Itoa(params.NbService) + "Service Created with each rate:" + params.ServRate + postResponse
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

func getCommonPorts(port_src []string, port_dst []string, number_of_service int) []string {
	common := []string{}

	// Use a map to store the count of each element in port_src
	count := make(map[string]int)
	for _, item := range port_src {
		count[item]++
	}

	// Iterate through port_dst and check if the item exists in port_src
	for _, item := range port_dst {
		if count[item] > 0 {
			common = append(common, item)
			if len(common) == number_of_service {
				break
			}
		}
	}

	return common
}

func getPortIDs(common_ports []string, BaseUrl string, NE_SRC string, NE_DST string) [][]string {

	fmt.Println("Get by port")

	var portPairs [][]string

	for _, X := range common_ports {
		// Build the URL strings for Src and Dst
		urlStr_Src := BaseUrl + "onc/ltp?name==" + X + "&ne.name==" + NE_SRC + "&select(id)"
		urlStr_Dst := BaseUrl + "onc/ltp?name==" + X + "&ne.name==" + NE_DST + "&select(id)"
		Port_Src_ID, err := client.GET(urlStr_Src)
		if err != nil {
			fmt.Println("Error getting client ID:", err)
			return nil // Return an empty string in case of an error
		}
		Port_Dst_ID, err := client.GET(urlStr_Dst)
		if err != nil {
			fmt.Println("Error getting client ID:", err)
			return nil // Return an empty string in case of an error
		}

		// Create a pair and append it to the array
		portPair := []string{Port_Src_ID, Port_Dst_ID}
		portPairs = append(portPairs, portPair)

	}

	return portPairs

}

func postPortIDs(Port_IDs [][]string, params ServiceParams) (string, error) {

	var postResponse string
	postUrlStr := params.BaseUrl + "onc/connection"
	for i := 0; i < params.NbService; i++ {
		service_name := "Service_" + strconv.Itoa(i+1)
		ServiceInfo := "posting to " + postUrlStr + " from source port ID: " + Port_IDs[i][0] + " to destination port ID: " + Port_IDs[i][1] + " with service name: " + service_name
		fmt.Println(ServiceInfo)
		postResponse, err := client.POST(postUrlStr, Port_IDs[i][0], Port_IDs[i][1], service_name, "ConnLpEthCbr", params.ServRate, "service")
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
