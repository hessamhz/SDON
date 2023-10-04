package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"napcore/internal/utils"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func CreateLP(urlStr string) {

	requestData := map[string]interface{}{
		"className": "Connection",
		"connEndPoints": []map[string]interface{}{
			{
				"className": "ConnEndPoint",
				"ltp": map[string]interface{}{
					"className": "Ltp",
					"id":        "ID1",
				},
				"endType": "source",
			},
			{
				"className": "ConnEndPoint",
				"ltp": map[string]interface{}{
					"className": "Ltp",
					"id":        "ID2",
				},
				"endType": "sink",
			},
		},
		"routingCriteria": "byLength",
		"sncpInfo": map[string]interface{}{
			"holdOffTime": 0,
			"revertive":   true,
			"sncpType":    "sncp_i",
			"wtrTime":     300,
		},
		"connLps": []map[string]interface{}{
			{
				"className": "ConnLpOtu",
				"rate":      "otu2x",
			},
		},
		"configurationState": "implemented",
		"hierarchicalLevel":  "infrastructure",
		"name":               "MyLightpath10G",
		"protection":         false,
	}

	// Convert the Go struct to a JSON string
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	// Create a new cookie jar
	jar, _ := cookiejar.New(nil)

	// Define the URL
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// Parse the cookies from your cookie file and add them to the cookie jar
	utils.ParseCookies(jar, "cookie.curl", parsedURL)

	// Create a custom HTTP client with insecure transport (bypassing SSL certificate verification)
	client := &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Create a GET request
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the JSON response
	fmt.Println(string(body))
}
