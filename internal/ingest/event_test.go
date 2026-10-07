package ingest

import (
	"encoding/json"
	"testing"
)

// The payload below is the contract from RFC §4.1; wh-backend must keep emitting it.
const samplePayload = `{
  "orderId": "o-0421",
  "userId": "u-operator-3",
  "destinationArea": "Sucursal Pocitos",
  "items": [
    { "sku": "TEC-002", "productId": "p-TEC-002", "quantity": 2 },
    { "sku": "TEC-004", "productId": "p-TEC-004", "quantity": 1 }
  ],
  "completedAt": "2026-08-30T14:22:05Z"
}`

func TestOrderCompletedDecodesContract(t *testing.T) {
	var ev OrderCompleted
	if err := json.Unmarshal([]byte(samplePayload), &ev); err != nil {
		t.Fatal(err)
	}
	if ev.OrderID != "o-0421" || len(ev.Items) != 2 || ev.Items[0].Quantity != 2 {
		t.Fatalf("decoded wrongly: %+v", ev)
	}
	if ev.CompletedAt.Year() != 2026 {
		t.Fatalf("completedAt not parsed: %v", ev.CompletedAt)
	}
}
