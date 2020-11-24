package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/alertmanager/template"
)

func webhookReceiver(rw http.ResponseWriter, rq *http.Request) {
	defer rq.Body.Close()
	data := template.Data{}
	err := json.NewDecoder(rq.Body).Decode(&data)

	if err != nil {
		panic(err)
	}

	fmt.Println("Status: ", data.Status)
	fmt.Println("Alerts:")

	for _, alert := range data.Alerts {
		fmt.Println("Status: ", alert.Status)
		fmt.Println("Summary: ", alert.Annotations["summary"])
		fmt.Println("Description: ", alert.Annotations["description"])
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(webhookReceiver)))
}
