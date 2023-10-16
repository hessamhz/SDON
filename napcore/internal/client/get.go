package client

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"napcore/internal/utils"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

// Modify the GET function to return the response body (ID)
func GET(urlStr string) (string, error) {
	// Create a new cookie jar
	jar, _ := cookiejar.New(nil)

	// Define the URL
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return "", err
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
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	result := strings.TrimSpace(strings.Trim(string(body), "[]"))

	// Return the response body as the ID
	return result, nil
}
