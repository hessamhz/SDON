package client

import (
	"napcore/internal/utils"

	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func DELETE(urlStr string) (bool, error) {
	// Create a new cookie jar
	jar, _ := cookiejar.New(nil)

	// Define the URL
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return false, err
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
	req, err := http.NewRequest("DELETE", urlStr, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false, err
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return false, err
	}

	// Print the JSON response
	fmt.Println("delete response:", string(body))

	return true, nil
}
