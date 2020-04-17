package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	t38c "github.com/lostpeer/tile38-client"
)

func printJSON(msg string, data interface{}) {
	b, _ := json.Marshal(data)
	fmt.Printf("%s: %s\n", msg, b)
}

func main() {
	tile38, err := t38c.New("localhost:9851", true)
	if err != nil {
		log.Fatal(err)
	}

	geofenceRequest := t38c.GeofenceNearby("buses", 33.5123, -112.2693, 200).
		Actions(t38c.Enter, t38c.Exit)

	busChan := t38c.NewChan("busstop", geofenceRequest)

	if err := tile38.SetChan(busChan); err != nil {
		log.Fatal(err)
	}

	events, err := tile38.Subscribe(context.Background(), "busstop")
	if err != nil {
		log.Fatal(err)
	}

	for event := range events {
		printJSON("event", event)
	}
}