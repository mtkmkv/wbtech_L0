package database

import (
	"L0/internal/models"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

// const values for connection to database
const (
	username = "postgres"
	password = "admin"
	host     = "localhost"
	port     = "5432"
	database = "db_wb_L0"
)

type PgxIface interface {
	Begin(context.Context) (pgx.Tx, error)
	Close(context.Context) error
}

// Connection - connect to database
func Connection() *pgx.Conn {
	connectionUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, database)
	conn, err := pgx.Connect(context.Background(), connectionUrl)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return nil
	}
	return conn
}

// AddMessageToDatabase - addition message to four tables of database
func AddMessageToDatabase(db PgxIface, order models.Order) error {
	ctx := context.Background()

	tx, err := db.Begin(ctx)
	defer db.Close(ctx)

	// add data to orders
	insertItemQuery := "INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shred) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);"
	_, err = tx.Exec(ctx, insertItemQuery,
		order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId,
		order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated, order.OofShred)
	if err != nil {
		tx.Rollback(ctx)
		return errors.New(fmt.Sprintf("Orders insertion failed (%v)\n", err))
	}

	// add data to delivery
	insertItemQuery = "INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
	_, err = tx.Exec(ctx, insertItemQuery,
		order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip,
		order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		tx.Rollback(ctx)
		return errors.New(fmt.Sprintf("Delivery insertion failed (%v)\n", err))
	}

	// add data to payment
	insertItemQuery = "INSERT INTO payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);"
	_, err = tx.Exec(ctx, insertItemQuery,
		order.OrderUID, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider,
		order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal,
		order.Payment.CustomFee)
	if err != nil {
		tx.Rollback(ctx)
		return errors.New(fmt.Sprintf("Payment insertion failed (%v)\n", err))
	}

	// add all items to database
	insertItemQuery = "INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);"
	for i := range order.Items {
		_, err = tx.Exec(ctx, insertItemQuery,
			order.OrderUID, order.Items[i].ChrtID, order.Items[i].TrackNumber, order.Items[i].Price, order.Items[i].Rid,
			order.Items[i].Name, order.Items[i].Sale, order.Items[i].Size, order.Items[i].TotalPrice, order.Items[i].NmID,
			order.Items[i].Brand, order.Items[i].Status)
		if err != nil {
			tx.Rollback(ctx)
			return errors.New(fmt.Sprintf("Item %v insertion failed (%v)\n", i, err))
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errors.New(fmt.Sprintf("Commit error: (%v)\n", err))
	} else {
		log.Printf("New message received.")
		return nil
	}
}

// SyncCacheAndDatabase - synchronize cache and database values
// copy values from database to cache
func SyncCacheAndDatabase(db PgxIface) error {
	//conn := connection()
	//defer closeConnection(conn)

	ctx := context.Background()

	tx, err := db.Begin(ctx)
	defer db.Close(ctx)

	var countOfRowsInTable int

	err = tx.QueryRow(ctx, "SELECT COUNT(*) FROM orders;").Scan(&countOfRowsInTable)
	if err != nil {
		return errors.New(fmt.Sprintf("QueryRow failed (%v)\n", err))
	}

	if len(models.Cache) != countOfRowsInTable {

		// copy orders to cache
		rows, err := tx.Query(ctx, "select * from orders;")
		if err != nil {
			return errors.New(fmt.Sprintf("QueryRow (orders) failed (%v)\n", err))
		}
		for rows.Next() {
			var order models.Order
			err := rows.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry,
				&order.Locale, &order.InternalSignature, &order.CustomerId,
				&order.DeliveryService, &order.Shardkey, &order.SmId,
				&order.DateCreated, &order.OofShred)
			if err != nil {
				return errors.New(fmt.Sprintf("Error in scanning order row (%v)\n", err))
			}
			if _, found := models.Cache[order.OrderUID]; !found {
				models.Cache[order.OrderUID] = order
			}
		}

		// copy delivery to cache
		rows, err = tx.Query(ctx, "SELECT * FROM delivery;")
		if err != nil {
			return errors.New(fmt.Sprintf("QueryRow (delivery) failed (%v)\n", err))
		}
		for rows.Next() {
			var delivery models.Delivery
			var uid string
			err := rows.Scan(&uid, &delivery.Name, &delivery.Phone,
				&delivery.Zip, &delivery.City, &delivery.Address, &delivery.Region,
				&delivery.Email)
			if err != nil {
				return errors.New(fmt.Sprintf("Error in scanning delivery row (%v)\n", err))
			}
			if value, found := models.Cache[uid]; found {
				value.Delivery = delivery
				models.Cache[value.OrderUID] = value
			}
		}

		// copy payment to cache
		rows, err = tx.Query(ctx, "SELECT * FROM payment;")
		if err != nil {
			return errors.New(fmt.Sprintf("QueryRow (payment) failed (%v)\n", err))
		}
		for rows.Next() {
			var payment models.Payment
			var uid string
			err := rows.Scan(&uid, &payment.Transaction, &payment.RequestID,
				&payment.Currency, &payment.Provider, &payment.Amount, &payment.PaymentDT,
				&payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee)
			if err != nil {
				return errors.New(fmt.Sprintf("Error in scanning payment row (%v)\n", err))
			}
			if value, found := models.Cache[uid]; found {
				value.Payment = payment
				models.Cache[value.OrderUID] = value
			}
		}

		// copy items to cache
		rows, err = tx.Query(ctx, "SELECT * FROM items;")
		if err != nil {
			return errors.New(fmt.Sprintf("QueryRow (items) failed (%v)\n", err))
		}
		for rows.Next() {
			var item models.Item
			var uid string
			err := rows.Scan(&uid, &item.ChrtID, &item.TrackNumber, &item.Price,
				&item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID,
				&item.Brand, &item.Status)
			if err != nil {
				return errors.New(fmt.Sprintf("Error in scanning item row (%v)\n", err))
			}
			if value, found := models.Cache[uid]; found {
				value.Items = append(value.Items, item)
				models.Cache[value.OrderUID] = value
			}
		}
	}

	log.Printf("The database and cache are synchronized.\n")
	return nil
}
