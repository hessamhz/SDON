package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"crypto/tls"
)

func main() {
	// Create a new cookie jar
	jar, _ := cookiejar.New(nil)

	// Define the URL
	// urlStr := "https://10.79.23.42/onc/ltp?name==OTU2x-1-1-1&ne.name==team1-NE-1&select(id)"
	urlStr := "https://10.79.23.42/onc/ltp?name==OTU2x-1-1-1&ne.name==team1-NE-2&select(id)"
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// Parse the cookies from your cookie file and add them to the cookie jar
	parseCookies(jar, "cookie.curl", parsedURL)

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
		return
	}

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

func parseCookies(jar *cookiejar.Jar, cookieFile string, url *url.URL) {
	// Read the cookie content from the file
	cookieContent, err := ioutil.ReadFile(cookieFile)
	if err != nil {
		fmt.Println("Error reading cookie file:", err)
		return
	}

	// Split the cookie file content into individual cookie lines
	cookieLines := strings.Split(string(cookieContent), "\n")

	// Parse and add cookies to the cookie jar
	for _, line := range cookieLines {
		fields := strings.Fields(line)
		if len(fields) >= 7 {
			domain := fields[0]
			httpOnly := fields[1] == "TRUE"
			path := fields[2]
			secure := fields[3] == "TRUE"
			// expire := fields[4]  // Commented out as it's not used
			name := fields[5]
			value := fields[6]

			// Create a new cookie
			cookie := &http.Cookie{
				Name:     name,
				Value:    value,
				Domain:   domain,
				Path:     path,
				Secure:   secure,
				HttpOnly: httpOnly,
			}

			// Add the cookie to the cookie jar
			jar.SetCookies(url, []*http.Cookie{cookie})
		}
	}
}
