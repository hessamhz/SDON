package functions

import (
	"encoding/json"
	"fmt"
	"napcore/internal/client"
	"strings"
)

type VisInfrastructureParameters struct {
	BaseUrl string
}

type Infrastructure struct {
	ConnName          string
	ConnId            int
	HierarchicalLevel string
	ConfigState       string
	LayerRate         string
	NeName1           string
	NeName2           string
	Port1             string
	Port2             string
	CreationTime      string
	ModificationTime  string
}

type InfrastructureResponse struct {
	Name               string `json:"name"`
	Id                 int    `json:"id"`
	HierarchicalLevel  string `json:"hierarchicalLevel"`
	ConfigurationState string `json:"configurationState"`
	TopmostLayerRate   string `json:"topmostLayerRate"`
	ConnEndPoints      []struct {
		LTP struct {
			NE struct {
				Name string `json:"name"`
			} `json:"ne"`
			Name string `json:"name"`
		} `json:"ltp"`
	} `json:"connEndPoints"`
	Log struct {
		CreationTime     string `json:"creationTime"`
		ModificationTime string `json:"modificationTime"`
	} `json:"log"`
}

func VisInfrastructure(params VisInfrastructureParameters) ([]Infrastructure, error) {
	urlInfra := params.BaseUrl + "onc/connection?hierarchicalLevel==infrastructure&view(connEndPoints.ltp.ne)"

	responseBody, err := client.GET(urlInfra)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil, err
	}
	//fmt.Println(responseBody)
	// Fix for concatenated JSON objects
	fixedResponseBody := "[" + strings.Join(strings.Split(responseBody, "}{"), "},{") + "]"

	var infraResponses []InfrastructureResponse
	if err := json.Unmarshal([]byte(fixedResponseBody), &infraResponses); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	var result []Infrastructure
	for _, infra := range infraResponses {
		res := Infrastructure{
			ConnName:          infra.Name,
			ConnId:            infra.Id,
			HierarchicalLevel: infra.HierarchicalLevel,
			ConfigState:       infra.ConfigurationState,
			LayerRate:         infra.TopmostLayerRate,
			NeName1:           infra.ConnEndPoints[0].LTP.NE.Name,
			NeName2:           infra.ConnEndPoints[1].LTP.NE.Name,
			Port1:             infra.ConnEndPoints[0].LTP.Name,
			Port2:             infra.ConnEndPoints[1].LTP.Name,
			CreationTime:      infra.Log.CreationTime,
			ModificationTime:  infra.Log.ModificationTime,
		}
		result = append(result, res)
	}

	return result, nil
}
