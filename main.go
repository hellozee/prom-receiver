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

	fmt.Println(data)
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(webhookReceiver)))
}
