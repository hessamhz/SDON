package main

import (
	"fmt"

	"napcore/env"
	"napcore/internal/functions"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

func main() {
	godotenv.Load()

	env := env.SetEnv()

	BASE_URL := env.BaseURL

	InfluxUrl := env.InfluxURL
	InfluxToken := env.InfluxToken
	InfluxBucket := env.InfluxBucket
	InfluxOrg := env.InfluxOrg

	err := runBashScript("update_cookies.sh")
	if err != nil {
		fmt.Println("Error executing bash script: ", err)
		return
	}
	writeToInfluxDBAsync(InfluxUrl, BASE_URL, InfluxToken, InfluxBucket, InfluxOrg)

	// Connect to NATS

	nc, err := nats.Connect("127.0.0.1:4222")
	if err != nil {
		fmt.Println("Error connecting to NATS:", err)
		return
	}
	defer nc.Close()

	// // Assuming BASE_URL is defined somewhere in your main or loaded from env
	// go publishStatusRegularly(nc, "visualize.service", getVisServiceData, BASE_URL)
	// go publishStatusRegularly(nc, "visualize.infrastructure", getVisInfraData, BASE_URL)
	// fmt.Println(getVisServiceData(BASE_URL))
	// fmt.Println(getVisInfraData(BASE_URL))
	// Listen for messages from NATS the topic "create.infrastructure"
	// Example of adding logging around subscription setup
	subscribeAndLog := func(subject string, handler func(*nats.Msg)) {
		_, err := nc.Subscribe(subject, func(msg *nats.Msg) {
			// Wrap the handler to add logging
			fmt.Println("Subscribing to subject: %s", subject)
			handler(msg)
		})
		if err != nil {
			fmt.Println("Error subscribing to subject %s: %v", subject, err)
		} else {
			fmt.Println("Successfully subscribed to subject: %s", subject)
		}
	}

	// Use the wrapper function for subscriptions
	subscribeAndLog("create.infrastructure", func(msg *nats.Msg) {
		handleNATSMessage(msg, BASE_URL)
	})

	subscribeAndLog("create.service", func(msg *nats.Msg) {
		handleNATSMessage(msg, BASE_URL)
	})

	subscribeAndLog("delete", func(msg *nats.Msg) {
		handleNATSMessage(msg, BASE_URL)
	})

	select {}

}
func writeToInfluxDBAsync(url, baseURL, token, bucket, org string) {
	go func() {
		functions.StartPeriodicDataWrite(url, baseURL, token, bucket, org)
	}()
}

func runBashScript(scriptPath string) error {
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdout = os.Stdout // Direct script's standard output to the program's output
	cmd.Stderr = os.Stderr // Direct script's standard error to the program's error output
	err := cmd.Run()
	return err
}

func handleNATSMessage(msg *nats.Msg, BASE_URL string) {
	// Handle the received message here
	fmt.Printf("Received message: %s\n", string(msg.Data))
	receivedMessage := string(msg.Data)
	values := strings.Split(receivedMessage, ",")
	fmt.Println(values)
	if values[0] == "CreateInfrastructure" {
		Params := functions.InfrastructureParams{
			BaseUrl:           BASE_URL,
			InfraLine:         "OTU2x-1-1-1",
			NeSrc:             values[1],
			NeDst:             values[2],
			ConnName:          values[3],
			HierarchicalLevel: "infrastructure",
		}
		fmt.Println("CreateInfra was called with Params:", Params)
		CreateInfraResponse, err := functions.CreateInfra(Params)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("CreateInfraResponse", CreateInfraResponse)
		}

	} else if values[0] == "CreateService" {
		intValue, _ := strconv.Atoi(values[4])

		ServiceParams := functions.ServiceParams{
			BaseUrl:   BASE_URL,
			NeSrc:     values[1],
			NeDst:     values[2],
			ServRate:  values[3],
			NbService: intValue,
		}

		fmt.Println("CreateLP was called with parameters ServiceParams", ServiceParams)
		CreateLPResponse := functions.CreateLP(ServiceParams)
		fmt.Println(CreateLPResponse)
	} else if values[0] == "DeleteService" {
		fmt.Println(values)
		// Handle the delete message
		connName := values[1]
		deleteInfrastructureAsWell := false
		DeleteConnResponse, err := functions.DeleteConn(BASE_URL, connName, deleteInfrastructureAsWell)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("DeleteConnResponse", DeleteConnResponse)
		}
	} else if values[0] == "DeleteInfrastructure" {
		fmt.Println(values)
		// Handle the delete message
		connName := values[1]
		deleteInfrastructureAsWell := true
		DeleteConnResponse, err := functions.DeleteConn(BASE_URL, connName, deleteInfrastructureAsWell)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("DeleteConnResponse", DeleteConnResponse)
		}
	} else if values[0] == "DeleteBoth" {
		fmt.Println(values)
		// Handle the delete message
		connNameService := values[2]
		deleteInfrastructureAsWell := false
		DeleteConnResponse, err := functions.DeleteConn(BASE_URL, connNameService, deleteInfrastructureAsWell)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("DeleteConnResponse", DeleteConnResponse)
		}
		connNameInfra := values[1]
		DeleteConnResponseNew, err := functions.DeleteConn(BASE_URL, connNameInfra, true)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("DeleteConnResponse", DeleteConnResponseNew)
		}

	}
}
