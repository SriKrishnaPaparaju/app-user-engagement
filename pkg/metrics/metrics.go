package metrics

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"gopkg.in/segmentio/analytics-go.v3"
	segmentAnalytics "gopkg.in/segmentio/analytics-go.v3"
)

// Struct to represent an instance of Metric
type Metric struct {
	Timestamp   string
	Name        string
	Description string
	Payload     string
}

func main() {
	fmt.Println(createTrackID())
	m := Metric{
		Timestamp:   "sasddad",
		Name:        "asdasdad",
		Description: "asdasfdadfs",
		Payload:     "adfafsasfsf",
	}
	fmt.Println(PublishMetricInstance(m))
}

// Create unique but random trackID for a user
func createTrackID() string {
	rand.Seed(time.Now().UnixNano())
	lengthOfRandomID := 16
	randomChars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789")
	var randomIDBuffer strings.Builder
	for index := 0; index < lengthOfRandomID; index++ {
		randomIDBuffer.WriteRune(randomChars[rand.Intn(len(randomChars))])
	}
	return randomIDBuffer.String()

}

// Receive metrics from application

// Publish metrics to backend storage (Segment for now)
func PublishMetricInstance(metricInstance Metric) bool {

	// Retrieve trackID
	// If trackID is not found, create a new trackID and store at the location indicated by FILEPATH_FOR_STORING_TRACKING / default to '/tmp'

	// Create JSON for the Metric Struct
	metricJSON, err := json.Marshal(metricInstance)
	if err != nil {
		fmt.Println(err)
		return false
	}
	// Publish Metrics
	return publishMetricsToBackendStore(string(metricJSON))
}

func publishMetricsToBackendStore(metricInstance string) bool {
	segmentClient := segmentAnalytics.New(readEnvironmentVariable("SEGMENT_TOKEN"))
	fmt.Println(readEnvironmentVariable("SEGMENT_TOKEN"))

	segementToken := readEnvironmentVariable("SEGMENT_SOURCE_NAME")
	fmt.Println(segementToken)

	defer segmentClient.Close()

	segmentClient.Enqueue(analytics.Track{
		UserId: segementToken,
		Event:  metricInstance,
		//		Properties: analytics.NewProperties().Set("plan", "trail"),
	})
	fmt.Println("sent")

	return true
}

// Read ENV
func readEnvironmentVariable(variableName string) string {
	variableValue := os.Getenv(variableName)
	return variableValue
}
