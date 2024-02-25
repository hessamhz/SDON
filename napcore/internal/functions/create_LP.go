package functions

import (
	"encoding/json"
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
	fmt.Println(postResponse)
	return strconv.Itoa(Params.NbService) + "Service Created with each rate:" + Params.ServRate
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
			if len(common) == 8 {
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

// Assuming ErrorResponse structure based on the JSON error message you provided.
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func postPortIDs(portIDs [][]string, Params ServiceParams) ([]string, error) {
	var responses []string
	postUrlStr := Params.BaseUrl + "onc/connection"
	fmt.Println("Port IDs:", portIDs)
	fmt.Println("Number of Services:", Params.NbService)

	for serviceIndex := 0; serviceIndex < Params.NbService; serviceIndex++ {
		serviceCreated := false

		for portIndex, portIDPair := range portIDs {
			serviceName := fmt.Sprintf("Service_%d_Attempt_%d", serviceIndex+1, portIndex+1)
			serviceInfo := fmt.Sprintf("Attempting to post to %s from source port ID: %s to destination port ID: %s with service name: %s", postUrlStr, portIDPair[0], portIDPair[1], serviceName)
			fmt.Println(serviceInfo)

			postResponse, err := client.POST(postUrlStr, portIDPair[0], portIDPair[1], serviceName, "ConnLpEthCbr", Params.ServRate, "service")
			if err != nil {
				// If there's an error, handle it according to its content
				fmt.Println("Error occurred:", err)
				continue // Proceed to the next port pair if error handling is not specified here
			}

			// Check if postResponse is an error response
			if isErrorResponse(postResponse) {
				fmt.Println("Detected Connection-endNotIdle error for", serviceName, "with port IDs", portIDPair)
				continue // Try next port pair for the same service
			} else {
				responses = append(responses, postResponse) // Successfully created the service
				serviceCreated = true
				break // Successfully created a service, proceed to the next service
			}
		}

		if !serviceCreated {
			fmt.Println("Failed to create service after trying all port pairs for service index", serviceIndex+1)
		}
	}

	return responses, nil // Return the successful responses
}

// Adjusted isErrorResponse to handle string that could be normal response or error
func isErrorResponse(response string) bool {
	var errResp ErrorResponse

	if err := json.Unmarshal([]byte(response), &errResp); err != nil {
		return false // If parsing fails, assume it's not the specific error we're looking for
	}

	return errResp.Status == "INTERNAL_SERVER_ERROR" && errResp.Message == "Connection-endNotIdle"
}
