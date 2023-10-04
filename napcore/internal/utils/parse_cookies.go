package utils

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)


func ParseCookies(jar *cookiejar.Jar, cookieFile string, url *url.URL) {
	// Open the cookie file
	file, err := os.Open(cookieFile)
	if err != nil {
		fmt.Println("Error opening cookie file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when done

	// Create a buffer to read the cookie content
	buf := make([]byte, 1024) // You can adjust the buffer size as needed

	// Read and parse cookies from the file
	for {
		n, err := file.Read(buf)
		if err != nil {
			break // End of file or error occurred
		}
		cookieContent := string(buf[:n])

		// Split the cookie file content into individual cookie lines
		cookieLines := strings.Split(cookieContent, "\n")

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
}
