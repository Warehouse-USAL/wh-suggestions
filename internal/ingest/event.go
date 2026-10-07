// Package ingest turns order.completed events into Redis updates.
package ingest

import "time"

// OrderCompleted is the payload wh-backend publishes on the order.completed topic.
type OrderCompleted struct {
	OrderID         string    `json:"orderId"`
	UserID          string    `json:"userId"`
	DestinationArea string    `json:"destinationArea"`
	Items           []Item    `json:"items"`
	CompletedAt     time.Time `json:"completedAt"`
}

// Item is one order line of an OrderCompleted event.
type Item struct {
	SKU       string `json:"sku"`
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
