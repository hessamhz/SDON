package abstractions

import (
	"fmt"
	"napcore/internal/client"
	"regexp"
	"strconv"
)

type ServiceParams struct {
	BASE_URL   string
	NE_SRC     string
	NE_DST     string
	NB_SERVICE int
}

func CreateLP(params ServiceParams) {
	// Get ports for Port_Src and Port_Dst
	Port_Src := getPorts(params.BASE_URL, params.NE_SRC)
	Port_Dst := getPorts(params.BASE_URL, params.NE_DST)
	fmt.Println("Port_SRC", Port_Src)
	fmt.Println("Port_DST", Port_Dst)

	Port_Common := getCommonPorts(Port_Src, Port_Dst, params.NB_SERVICE)
	fmt.Println("Common Ports", Port_Common)

	Port_IDs := getPortIDs(Port_Common, params.BASE_URL, params.NE_SRC, params.NE_DST)
	fmt.Println("Port_IDs:", Port_IDs)

	fmt.Println("Port_IDs:", Port_IDs)
	post_port_IDs(Port_IDs, params)

}

func getPorts(BASE_URL, NEName string) []string {
	fmt.Println("Ports")
	urlStr := BASE_URL + "onc/ltp?ltpType==physical&ne.name==" + NEName + "&select(id,name)&name==OGBE1-*"

	Ports, err := client.GET(urlStr)
	if err != nil {
		fmt.Println("Error getting ports:", err)
		return nil
	}
	fmt.Println("Fetched Ports:", Ports)
	// Define a regular expression pattern to match "name" values

	pattern := `"name":"([^"]+)"`

	// Compile the regular expression pattern
	re := regexp.MustCompile(pattern)

	// Find all matches in the input string
	matches := re.FindAllStringSubmatch(Ports, -1)

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

func getPortIDs(common_ports []string, BASE_URL string, NE_SRC string, NE_DST string) [][]string {

	fmt.Println("Get by port")

	var portPairs [][]string

	for _, X := range common_ports {
		// Build the URL strings for Src and Dst
		urlStr_Src := BASE_URL + "onc/ltp?name==" + X + "&ne.name==" + NE_SRC + "&select(id)"
		urlStr_Dst := BASE_URL + "onc/ltp?name==" + X + "&ne.name==" + NE_DST + "&select(id)"
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

func post_port_IDs(Port_IDs [][]string, params ServiceParams) {
	postUrlStr := params.BASE_URL + "onc/connection"
	for i := 0; i < params.NB_SERVICE; i++ {
		service_name := "Service_" + strconv.Itoa(i+1)
		what_iam_doing := "I am posting to " + postUrlStr + " from source port ID: " + Port_IDs[i][0] + " to destination port ID: " + Port_IDs[i][1] + " with service name: " + service_name
		fmt.Println(what_iam_doing)
		client.POST(postUrlStr, Port_IDs[i][0], Port_IDs[i][1], service_name, "ConnLpEthCbr", "1Gb", "service")

	}
	//client.POST(postUrlStr, Port_IDs[0][0], Port_IDs[0][1], "TestSecSrvc", "ConnLpEthCbr", "1Gb", "service")

}
