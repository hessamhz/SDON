package functions

import (
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// The entry function to start the periodic data fetch and write process
func StartPeriodicDataWrite(url string, baseUrl string, token string, bucket string, org string) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		write_to_influxdb(url, baseUrl, token, bucket, org)
	}
}

func write_to_influxdb(url string, baseUrl string, token string, bucket string, org string) {
	// Create client

	client := influxdb2.NewClient(url, token)
	defer client.Close()

	// Get a write API for the organization and bucket
	writeAPI := client.WriteAPI(org, bucket)

	serviceData := getVisServiceData(baseUrl)
	infraData := getVisInfraData(baseUrl)
	// Prepare write API
	fmt.Println(serviceData, infraData)

	// Write serviceData
	for _, service := range serviceData {
		p := influxdb2.NewPointWithMeasurement("serviceData").
			AddTag("ConnId", fmt.Sprintf("%d", service.ConnId)).
			AddField("ConnName", service.ConnName).
			AddField("HierarchicalLevel", service.HierarchicalLevel).
			AddField("ConfigState", service.ConfigState).
			AddField("LayerRate", service.LayerRate).
			AddField("NeName1", service.NeName1).
			AddField("NeName2", service.NeName2).
			AddField("Port1", service.Port1).
			AddField("Port2", service.Port2).
			AddField("CreationTime", service.CreationTime).
			AddField("ModificationTime", service.ModificationTime).
			SetTime(time.Now())
		writeAPI.WritePoint(p)
	}

	// Write infraData
	for _, infra := range infraData {
		p := influxdb2.NewPointWithMeasurement("infraData").
			AddTag("ConnId", fmt.Sprintf("%d", infra.ConnId)).
			AddField("ConnName", infra.ConnName).
			AddField("HierarchicalLevel", infra.HierarchicalLevel).
			AddField("ConfigState", infra.ConfigState).
			AddField("LayerRate", infra.LayerRate).
			AddField("NeName1", infra.NeName1).
			AddField("NeName2", infra.NeName2).
			AddField("Port1", infra.Port1).
			AddField("Port2", infra.Port2).
			AddField("CreationTime", infra.CreationTime).
			AddField("ModificationTime", infra.ModificationTime).
			SetTime(time.Now())
		writeAPI.WritePoint(p)
	}
	fmt.Println("added data into influxdb")

	// Assuming you're using writeAPI (non-blocking) from your example
	errorsCh := writeAPI.Errors()
	go func() {
		for err := range errorsCh {
			fmt.Printf("write error: %v\n", err)
		}
	}()
	// Ensure all buffered points are sent
	writeAPI.Flush()

	// Check for errors if necessary (not shown here)
	// Optionally, check for errors from the Flush operation
	// Check for errors if necessary

}

func getVisServiceData(baseUrl string) []Service {
	visParameters := VisParameters{BaseUrl: baseUrl}
	response, err := VisService(visParameters)
	if err != nil {
		fmt.Println("Error retrieving service visualization:", err)
		return nil // Return an empty slice or nil if there's an error
	}

	return response // Assuming response is of type []VisualizationData
}

func getVisInfraData(baseUrl string) []Infrastructure {
	visInfrastructureParameters := VisInfrastructureParameters{BaseUrl: baseUrl}
	response, err := VisInfrastructure(visInfrastructureParameters)
	if err != nil {
		fmt.Println("Error retrieving infrastructure visualization:", err)
		return nil // Return an empty slice or nil if there's an error
	}

	return response // Assuming response is of type []VisualizationData
}

/*type Infrastructure struct {
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

type Service struct {
	ConnName          string
	ConnId            int
	HierarchicalLevel string
	ConfigState       string
	LayerRate         string
	NeName1           string
	NeName2           string // Extracted manually
	Port1             string // Extracted manually
	Port2             string // Extracted manually
	CreationTime      string // Extracted manually
	ModificationTime  string // Extracted manually
}
*/
