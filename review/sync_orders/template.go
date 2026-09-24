package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Order struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func syncOrders(db *sql.DB, writer *kafka.Writer) {
	for {
		rows, err := db.Query("SELECT id, user_id, amount, created_at FROM orders WHERE synced = false")
		if err != nil {
			log.Printf("query error: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		var orders []Order
		for rows.Next() {
			var o Order
			if err := rows.Scan(&o.ID, &o.UserID, &o.Amount, &o.CreatedAt); err != nil {
				log.Printf("scan error: %v", err)
				continue
			}
			orders = append(orders, o)
		}
		rows.Close()

		for _, o := range orders {
			data, _ := json.Marshal(o)
			err := writer.WriteMessages(context.Background(), kafka.Message{Value: data})
			if err != nil {
				log.Printf("kafka write error: %v", err)
				continue
			}
			_, err = db.Exec("UPDATE orders SET synced = true WHERE id = ?", o.ID)
			if err != nil {
				log.Printf("update error: %v", err)
			}
		}

		time.Sleep(10 * time.Second)
	}
}
