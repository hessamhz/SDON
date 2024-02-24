package functions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nats-io/nats.go"
)

func handleNATSMessage(msg *nats.Msg, BASE_URL string) {
	// Handle the received message here
	fmt.Printf("Received message: %s\n", string(msg.Data))
	receivedMessage := string(msg.Data)
	values := strings.Split(receivedMessage, ",")
	if values[0] == "create.infrastructure" {
		Params := InfrastructureParams{
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
	} else if values[0] == "create.service" {
		intValue, _ := strconv.Atoi(values[4])

		ServiceParams := ServiceParams{
			BaseUrl:   BASE_URL,
			NeSrc:     values[1],
			NeDst:     values[2],
			ServRate:  values[3],
			NbService: intValue,
		}

		fmt.Println("CreateLP was called with parameters ServiceParams", ServiceParams)
		// CreateLPResponse := functions.CreateLP(ServiceParams)
	} else if values[0] == "delete" {
		// Handle the delete message
		// DeleteConnResponse, err := functions.DeleteConn(ServParams, InfParams, deleteInfrastructure)
		// if err != nil {
		// 	fmt.Println("Error", err)
		// } else {
		// 	fmt.Println("DeleteConnResponse", DeleteConnResponse)
		// }
	}
}
