package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/buger/jsonparser"
)

func GetSupportedApplications() (string, error) {
	supportedApplicationsFile, err := os.Open("supported-apps.json")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fileInBytes, _ := ioutil.ReadAll(supportedApplicationsFile)
	return string(fileInBytes), nil
}

func processSupportedApplications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	response, err := GetSupportedApplications()
	if err != nil {
		response = "Could not retrieve supported applications. Please try again later"
	}
	w.Write([]byte(response))
}

func processGetMetrics(w http.ResponseWriter, r *http.Request) {
	response := "Not implemented yet"
	w.Write([]byte(response))
}

func processPublishMetrics(w http.ResponseWriter, r *http.Request) {

	// Read body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Printf("body is %s : ", string(body))
	jsonparser.ArrayEach(body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		v, _, _, _ := jsonparser.Get(value, "application_name")
		fmt.Println(string(v))
	}, "Metrics", "metric_instances")

	/*
		a1, a2, a3, a4 := jsonparser.Get(body, "Metrics", "metric_instances", "[0]", "application_name")
		fmt.Printf("a1 is %s : ", a1)
		fmt.Println("")
		fmt.Printf("a2 is %s : ", a2)
		fmt.Println("")
		fmt.Printf("a3 is %s : ", a3)
		fmt.Println("")
		fmt.Printf("a4 is %s : ", a4)
		fmt.Println("")
	*/
	w.Write([]byte("something"))

}

func main() {
	http.HandleFunc("/api/app_usage_metrics/v1/applications", processSupportedApplications)
	http.HandleFunc("/api/app_usage_metrics/v1/metrics", processGetMetrics)
	http.HandleFunc("/api/app_usage_metrics/v1/metrics/publish", processPublishMetrics)
	address := ":9000"
	log.Println("Starting server on address", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
