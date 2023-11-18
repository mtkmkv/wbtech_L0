package database

import (
	"L0/internal/models"
	"context"
	"errors"
	"testing"

	"github.com/pashagolub/pgxmock/v2"
)

type mockBehavior func(order models.Order)

func TestAddMessageToDatabase(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close(context.Background())

	tests := []struct {
		name    string
		mock    mockBehavior
		input   models.Order
		wantErr bool
	}{
		{
			name: "ok",
			input: models.Order{
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
					PaymentDT:    1817,
					Amount:       1637907727,
					Bank:         "alpha",
					DeliveryCost: 1500,
					GoodsTotal:   317,
					CustomFee:    0,
				},
				Items: []models.Item{{
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
				}},
				Locale:            "en",
				InternalSignature: "",
				CustomerId:        "test",
				DeliveryService:   "meest",
				Shardkey:          "9",
				SmId:              99,
				DateCreated:       "2021-11-26T06:22:19Z",
				OofShred:          "1",
			},
			mock: func(order models.Order) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO orders").
					WithArgs(order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId,
						order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated, order.OofShred).WillReturnResult(pgxmock.NewResult("INSERT", 1))

				mock.ExpectExec("INSERT INTO delivery").
					WithArgs(order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip,
						order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email).WillReturnResult(pgxmock.NewResult("INSERT", 1))

				mock.ExpectExec("INSERT INTO payment").
					WithArgs(order.OrderUID, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider,
						order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal,
						order.Payment.CustomFee).WillReturnResult(pgxmock.NewResult("INSERT", 1))

				for i := range order.Items {
					mock.ExpectExec("INSERT INTO items").
						WithArgs(order.OrderUID, order.Items[i].ChrtID, order.Items[i].TrackNumber, order.Items[i].Price, order.Items[i].Rid,
							order.Items[i].Name, order.Items[i].Sale, order.Items[i].Size, order.Items[i].TotalPrice, order.Items[i].NmID,
							order.Items[i].Brand, order.Items[i].Status).WillReturnResult(pgxmock.NewResult("INSERT", 1))
				}

				mock.ExpectCommit()
			},
		},
		{
			name: "rollback order",
			input: models.Order{
				OrderUID:    "9",
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
					PaymentDT:    1817,
					Amount:       16379077275,
					Bank:         "alpha",
					DeliveryCost: 1500,
					GoodsTotal:   317,
					CustomFee:    0,
				},
				Items: []models.Item{{
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
				}},
				Locale:            "en",
				InternalSignature: "",
				CustomerId:        "test",
				DeliveryService:   "meest",
				Shardkey:          "9",
				SmId:              99,
				DateCreated:       "2021-11-26T06:22:19Z",
				OofShred:          "1",
			},
			wantErr: true,
			mock: func(order models.Order) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO orders").
					WithArgs(order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId,
						order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated, order.OofShred).WillReturnError(errors.New("error"))

				mock.ExpectRollback()
			},
		},
		{
			name: "rollback delivery",
			input: models.Order{
				OrderUID:    "9",
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
					PaymentDT:    1817,
					Amount:       16379077275,
					Bank:         "alpha",
					DeliveryCost: 1500,
					GoodsTotal:   317,
					CustomFee:    0,
				},
				Items: []models.Item{{
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
				}},
				Locale:            "en",
				InternalSignature: "",
				CustomerId:        "test",
				DeliveryService:   "meest",
				Shardkey:          "9",
				SmId:              99,
				DateCreated:       "2021-11-26T06:22:19Z",
				OofShred:          "1",
			},
			wantErr: true,
			mock: func(order models.Order) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO orders").
					WithArgs(order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId,
						order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated, order.OofShred).WillReturnResult(pgxmock.NewResult("INSERT", 1))

				mock.ExpectExec("INSERT INTO delivery").
					WithArgs(order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip,
						order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email).WillReturnError(errors.New("error"))

				mock.ExpectRollback()
			},
		},
		{
			name: "rollback payment",
			input: models.Order{
				OrderUID:    "9",
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
					PaymentDT:    1817,
					Amount:       16379077275,
					Bank:         "alpha",
					DeliveryCost: 1500,
					GoodsTotal:   317,
					CustomFee:    0,
				},
				Items: []models.Item{{
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
				}},
				Locale:            "en",
				InternalSignature: "",
				CustomerId:        "test",
				DeliveryService:   "meest",
				Shardkey:          "9",
				SmId:              99,
				DateCreated:       "2021-11-26T06:22:19Z",
				OofShred:          "1",
			},
			wantErr: true,
			mock: func(order models.Order) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO orders").
					WithArgs(order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId,
						order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated, order.OofShred).WillReturnResult(pgxmock.NewResult("INSERT", 1))

				mock.ExpectExec("INSERT INTO delivery").
					WithArgs(order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip,
						order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email).WillReturnResult(pgxmock.NewResult("INSERT", 1))

				mock.ExpectExec("INSERT INTO payment").
					WithArgs(order.OrderUID, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider,
						order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal,
						order.Payment.CustomFee).WillReturnError(errors.New("error"))

				mock.ExpectRollback()
			},
		},
		{
			name: "rollback items",
			input: models.Order{
				OrderUID:    "9",
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
					PaymentDT:    1817,
					Amount:       16379077275,
					Bank:         "alpha",
					DeliveryCost: 1500,
					GoodsTotal:   317,
					CustomFee:    0,
				},
				Items: []models.Item{{
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
				}},
				Locale:            "en",
				InternalSignature: "",
				CustomerId:        "test",
				DeliveryService:   "meest",
				Shardkey:          "9",
				SmId:              99,
				DateCreated:       "2021-11-26T06:22:19Z",
				OofShred:          "1",
			},
			wantErr: true,
			mock: func(order models.Order) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO orders").
					WithArgs(order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId,
						order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated, order.OofShred).WillReturnResult(pgxmock.NewResult("INSERT", 1))

				mock.ExpectExec("INSERT INTO delivery").
					WithArgs(order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip,
						order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email).WillReturnResult(pgxmock.NewResult("INSERT", 1))

				mock.ExpectExec("INSERT INTO payment").
					WithArgs(order.OrderUID, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider,
						order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal,
						order.Payment.CustomFee).WillReturnResult(pgxmock.NewResult("INSERT", 1))

				for i := range order.Items {
					mock.ExpectExec("INSERT INTO items").
						WithArgs(order.OrderUID, order.Items[i].ChrtID, order.Items[i].TrackNumber, order.Items[i].Price, order.Items[i].Rid,
							order.Items[i].Name, order.Items[i].Sale, order.Items[i].Size, order.Items[i].TotalPrice, order.Items[i].NmID,
							order.Items[i].Brand, order.Items[i].Status).WillReturnError(errors.New("error"))
				}

				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)

			// now we execute our method
			err := AddMessageToDatabase(mock, tt.input)
			if tt.wantErr {
				t.Log(err)
			} else {
				t.Log("ok")
			}

			// we make sure that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
