package main

import (
	"fmt"
	"napcore/env"
	"napcore/internal/client"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	env := env.SetEnv()

	fmt.Println("Interfaces:")
	urlStr := env.BaseURL + "onc/ltp?name==OTU2x-1-1-1&ne.name==team1-NE-1&select(id)"
	client.GET(urlStr)
	urlStr = env.BaseURL + "onc/ltp?name==OTU2x-1-1-1&ne.name==team1-NE-2&select(id)"
	client.GET(urlStr)

	fmt.Println("Ports")
	urlStr = env.BaseURL + "onc/ltp?ltpType==physical&ne.name==team1-NE-1&select(id,name)&name==OGBE*" 
	client.GET(urlStr)
	urlStr = env.BaseURL + "onc/ltp?ltpType==physical&ne.name==team1-NE-2&select(id,name)&name==OGBE*"
	client.GET(urlStr)

	fmt.Println("Get by port")
	urlStr = env.BaseURL + "onc/ltp?name==OGBE1-1-1-12&ne.name==team1-NE-1&select(id)"
	client.GET(urlStr)
	urlStr = env.BaseURL + "onc/ltp?name==OGBE1-1-1-12&ne.name==team1-NE-2&select(id)"

	// postUrlStr := env.BaseURL + "onc/connection"
	//client.POST(postUrlStr, "94", "152", "TestSecSrvc", "ConnLpEthCbr", "1Gb", "service") // id: 85
	// client.POST(postUrlStr, "227", "257", "TestFirstConHessam", "ConnLpOtu", "otu2x", "infrastructure") // id:76

	fmt.Println("Get Service")
	urlStr = env.BaseURL + "onc/connection?name==FatihConnection&select(id)"
	client.GET(urlStr)

	fmt.Println("Deleting")
	urlStr = env.BaseURL + "onc/connection/78"
	client.DELETE(urlStr)

}

