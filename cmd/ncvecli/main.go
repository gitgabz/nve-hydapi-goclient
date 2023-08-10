package main

import (
	"gitgabz/nve-hydapi-goclient/internal/config"
	"gitgabz/nve-hydapi-goclient/internal/nveapi"
	"log"
)

func main() {

	client := nveapi.NewClient(config.ReturnApiKey())

	q := nveapi.RequestQueryObservations{
		StationId:      "12.534.0", // Mjøndalen bru
		Parameter:      "1000",     // Vannstand
		ResolutionTime: "inst",     // Siste målerstand
	}

	jresp, err := client.GetObservations(q)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	log.Printf("Response: %.3f%s", jresp.Data[0].Observations[0].Value, jresp.Data[0].Unit)

}
