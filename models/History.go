package models

import "time"

type History struct {
    ID        string    `json:"id"`
    CustomerID string    `json:"customer_id"`
    Action    string    `json:"action"`
    Amount    float64   `json:"amount,omitempty"` // Optional, hanya untuk transaksi
    Timestamp time.Time `json:"timestamp"`
}
