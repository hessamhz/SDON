package functions

import (
	"encoding/json"
	"fmt"
	"napcore/internal/client"
	"strings"
)

type VisParameters struct {
	BaseUrl string
}

type Service struct {
	ConnName          string
	ConnId            int
	HierarchicalLevel string
	ConfigState       string
	LayerRate         string
	NeName1           string // Extracted manually
	NeName2           string // Extracted manually
	Port1             string // Extracted manually
	Port2             string // Extracted manually
	CreationTime      string // Extracted manually
	ModificationTime  string // Extracted manually
}

// Assuming JSON structure for `connEndPoints` and `log`
type ConnEndPoint struct {
	LTP struct {
		NE struct {
			Name string `json:"name"`
		} `json:"ne"`
		Name string `json:"name"`
	} `json:"ltp"`
}

type JSONService struct {
	Name              string         `json:"name"`
	Id                int            `json:"id"`
	HierarchicalLevel string         `json:"hierarchicalLevel"`
	ConfigState       string         `json:"configurationState"`
	LayerRate         string         `json:"topmostLayerRate"`
	ConnEndPoints     []ConnEndPoint `json:"connEndPoints"`
	Log               struct {
		CreationTime     string `json:"creationTime"`
		ModificationTime string `json:"modificationTime"`
	} `json:"log"`
}

func VisService(params VisParameters) ([]Service, error) {
	urlService := params.BaseUrl + "onc/connection?hierarchicalLevel==service&view(connEndPoints.ltp.ne)"

	responseBody, err := client.GET(urlService)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil, err
	}
	//fmt.Println(responseBody)
	// wrapping it in an array if necessary
	fixedResponseBody := "[" + strings.Join(strings.Split(responseBody, "},"), "},") + "]"
	var jsonServices []JSONService
	if err := json.Unmarshal([]byte(fixedResponseBody), &jsonServices); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	var services []Service
	for _, js := range jsonServices {
		service := Service{
			ConnName:          js.Name,
			ConnId:            js.Id,
			HierarchicalLevel: js.HierarchicalLevel,
			ConfigState:       js.ConfigState,
			LayerRate:         js.LayerRate,
		}

		if len(js.ConnEndPoints) > 1 {
			service.NeName1 = js.ConnEndPoints[0].LTP.NE.Name
			service.NeName2 = js.ConnEndPoints[1].LTP.NE.Name
			service.Port1 = js.ConnEndPoints[0].LTP.Name
			service.Port2 = js.ConnEndPoints[1].LTP.Name
		}

		service.CreationTime = js.Log.CreationTime
		service.ModificationTime = js.Log.ModificationTime

		services = append(services, service)
	}

	return services, nil
}
