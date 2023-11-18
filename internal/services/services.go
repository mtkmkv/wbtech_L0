package services

import (
	"L0/internal/database"
	"L0/internal/models"
	"encoding/json"
	"log"

	"github.com/nats-io/stan.go"
)

func unmarshalMessage(m []byte) (models.Order, error) {
	var order models.Order

	err := json.Unmarshal(m, &order)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

// GetMessage from the channel
func GetMessage() {

	sc, _ := stan.Connect("test-cluster", "sub1")

	_, err := sc.Subscribe("msg", func(m *stan.Msg) {
		order, err := unmarshalMessage(m.Data)
		if err != nil {
			log.Printf("Error in marshaling message (incorrect messege type): %v\n", err)
		} else {
			err = database.AddMessageToDatabase(database.Connection(), order) // add to database
			if err != nil {
				log.Print(err)
			} else {
				models.Cache[order.OrderUID] = order // add to cache
			}
		}
	})
	if err != nil {
		log.Printf("Error in subscription: %v\n", err)
	}
}
