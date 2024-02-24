package main

import (
	"encoding/json"
	"fmt"

	"napcore/env"
	"napcore/internal/functions"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

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
	functions.StartPeriodicDataWrite(InfluxUrl, BASE_URL, InfluxToken, InfluxBucket, InfluxOrg)

	// Connect to NATS
	/*
		nc, err := nats.Connect(nats.DefaultURL)
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
		var e error
		_, e = nc.Subscribe("create.infrastructure", func(msg *nats.Msg) {
			handleNATSMessage(msg, BASE_URL)
		})

		if e != nil {
			fmt.Println("Error subscribing to NATS:", err)
			return
		}

		_, e = nc.Subscribe("create.service", func(msg *nats.Msg) {
			handleNATSMessage(msg, BASE_URL)
		})

		if e != nil {
			fmt.Println("Error subscribing to NATS:", err)
			return
		}

		_, e = nc.Subscribe("delete", func(msg *nats.Msg) {
			handleNATSMessage(msg, BASE_URL)
		})

		if e != nil {
			fmt.Println("Error subscribing to NATS:", err)
			return
		}

		// Listen for messages from NATS the topic "delete"
		// _, e = nc.Subscribe("delete", func(msg *nats.Msg) {
		// 	// Handle the received message here
		// 	fmt.Printf("Received message: %s\n", string(msg.Data))
		// 	receivedMessage := string(msg.Data)
		// 	values := strings.Split(receivedMessage, ",")
		// 	var delete, neSrc, neDst, connName, deleteInfrastructureAsWell string
		// 	if len(values) >= 2 {
		// 		delete = values[0]
		// 		neSrc = values[1]
		// 		neDst = values[2]
		// 		connName = values[3]
		// 		deleteInfrastructureAsWell = values[4]
		// 	}

		// 	fmt.Println("delete:", delete)
		// 	fmt.Println("neSrc:", neSrc)
		// 	fmt.Println("neDst:", neDst)
		// 	fmt.Println("connName:", connName)
		// 	fmt.Println("deleteInfrastructureAsWell:", deleteInfrastructureAsWell)
		// })
		// if e != nil {
		// 	fmt.Println("Error subscribing to NATS:", err)
		// 	return
		// }
		select {}
		/*
				Params := functions.InfrastructureParams{
					BaseUrl:           BASE_URL,
					InfraLine:         "OTU2x-1-1-1",
					NeSrc:             "team1-NE-1",
					NeDst:             "team1-NE-2",
					ConnName:          "FatihConnection",
					HierarchicalLevel: "infrastructure",
				}

				//Only  one 10GB or 	up to 8 1GB Service is allowed

				ServiceParams := functions.ServiceParams{
					BaseUrl:   BASE_URL,
					NeSrc:     "team1-NE-1",
					NeDst:     "team1-NE-2",
					ServRate:  "10Gb",
					NbService: 1,
				}

			VisParameters := functions.VisParameters{
				BaseUrl: BASE_URL,
			}
			VisInfrastructureParameters := functions.VisInfrastructureParameters{
				BaseUrl: BASE_URL,
			}

			/*
				DeleteInfrastructureAsWell := true

				CreateInfraResponse, err := functions.CreateInfra(Params)
				if err != nil {
					fmt.Println("Error", err)
				} else {
					fmt.Println("CreateInfraResponse", CreateInfraResponse)
				}

				fmt.Println("Infrastructure Created now its turn for CreateLPResponse ")

				CreateLPResponse := functions.CreateLP(ServiceParams)
				fmt.Println("CreateLPResponse", CreateLPResponse)
				fmt.Println("Lightpath Created now its turn for Delete Connections ")

				DeleteResponse, err := functions.DeleteConn(ServiceParams, Params, DeleteInfrastructureAsWell)
				if err != nil {
					fmt.Println("Error", err)
				} else {
					fmt.Println("DeleteResponse", DeleteResponse)
				}

			VisServiceResponse, err := functions.VisService(VisParameters)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				fmt.Println("CurrentServices", VisServiceResponse)
			}
			VisInfraResponse, err := functions.VisInfrastructure(VisInfrastructureParameters)
			if err != nil {
				fmt.Println("Error", err)
			} else {
				fmt.Println("CurrentInfrastructures", VisInfraResponse)
			}
	*/
}

func runBashScript(scriptPath string) error {
	cmd := exec.Command("bash", scriptPath)
	cmd.Stdout = os.Stdout // Direct script's standard output to the program's output
	cmd.Stderr = os.Stderr // Direct script's standard error to the program's error output
	err := cmd.Run()
	return err
}

func publishStatusRegularly[T any](nc *nats.Conn, subject string, dataFunc func(string) []T, baseUrl string) {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		data := dataFunc(baseUrl)
		if len(data) == 0 {
			fmt.Println("Skipping publish due to empty data")
			continue
		}
		jsonData, err := json.Marshal(data) // Serialize data to JSON
		if err != nil {
			fmt.Printf("Error serializing data for subject %s: %v\n", subject, err)
			continue
		}
		if err := nc.Publish(subject, jsonData); err != nil {
			fmt.Printf("Error publishing to subject %s: %v\n", subject, err)
		} else {
			fmt.Printf("Published to subject %s\n", subject)
		}
	}
}

func getVisServiceData(baseUrl string) []functions.Service {
	visParameters := functions.VisParameters{BaseUrl: baseUrl}
	response, err := functions.VisService(visParameters)
	if err != nil {
		fmt.Println("Error retrieving service visualization:", err)
		return nil // Return an empty slice or nil if there's an error
	}

	return response // Assuming response is of type []VisualizationData
}

func getVisInfraData(baseUrl string) []functions.Infrastructure {
	visInfrastructureParameters := functions.VisInfrastructureParameters{BaseUrl: baseUrl}
	response, err := functions.VisInfrastructure(visInfrastructureParameters)
	if err != nil {
		fmt.Println("Error retrieving infrastructure visualization:", err)
		return nil // Return an empty slice or nil if there's an error
	}

	return response // Assuming response is of type []VisualizationData
}

func handleNATSMessage(msg *nats.Msg, BASE_URL string) {
	// Handle the received message here
	fmt.Printf("Received message: %s\n", string(msg.Data))
	receivedMessage := string(msg.Data)
	values := strings.Split(receivedMessage, ",")
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
		// CreateInfraResponse, err := functions.CreateInfra(Params)
		// if err != nil {
		// 	fmt.Println("Error", err)
		// } else {
		// 	fmt.Println("CreateInfraResponse", CreateInfraResponse)
		// }

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
		// CreateLPResponse := functions.CreateLP(ServiceParams)

	} else if values[0] == "DeleteInfrastructure" {
		neSrc := values[1]
		neDst := values[2]
		connName := values[3]
		deleteInfrastructureAsWell := values[4]

		fmt.Println("neSrc:", neSrc)
		fmt.Println("neDst:", neDst)
		fmt.Println("connName:", connName)
		fmt.Println("deleteInfrastructureAsWell:", deleteInfrastructureAsWell)

		// Handle the delete message
		// DeleteConnResponse, err := functions.DeleteConn(ServParams, InfParams, deleteInfrastructure)
		// if err != nil {
		// 	fmt.Println("Error", err)
		// } else {
		// 	fmt.Println("DeleteConnResponse", DeleteConnResponse)
		// }
	}
}
