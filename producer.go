package main

import (
	"L0/internal/models"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, _ := stan.Connect("test-cluster", "simple-pub")
	defer sc.Close()

	item := models.Item{
		ChrtID:      9934930,
		TrackNumber: "WBILMTESTTRACK",
		Price:       453,
		Rid:         "ab4219087a764ae0btest",
		Name:        "Mascaras",
		Sale:        30,
		Size:        "0",
		TotalPrice:  317,
		NmID:        2389212,
		Brand:       "Vivienne Sabo",
		Status:      202,
	}

	var items []models.Item
	items = append(items, item)

	order := models.Order{
		OrderUID:    "b563feb7b2b84b6test",
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: models.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: models.Payment{
			Transaction:  "b563feb7b2b84b6test",
			RequestID:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       1817,
			PaymentDT:    1637907727,
			Bank:         "alpha",
			DeliveryCost: 1500,
			GoodsTotal:   317,
			CustomFee:    0,
		},
		Items:             items,
		Locale:            "en",
		InternalSignature: "",
		CustomerId:        "test",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmId:              99,
		DateCreated:       "2021-11-26T06:22:19Z",
		OofShred:          "1",
	}

	for i := 3; ; i++ {

		order.OrderUID = strconv.Itoa(i)            // create unique identifier
		order.Payment.Transaction = strconv.Itoa(i) // create unique identifier
		order.Items[0].Price = int64(1000 + i)      // create unique identifier
		order.Payment.Amount = int64(i * i * 100)
		order.Delivery.Address = "Ploshad Mira " + strconv.Itoa(i+10)
		record, _ := json.Marshal(order)

		err := sc.Publish("msg", record)
		if err != nil {
			log.Printf("Error in publishing message: %v\n", err)
		}

		time.Sleep(15 * time.Second)
	}
}
