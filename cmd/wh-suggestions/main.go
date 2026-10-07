// Command wh-suggestions serves order-history based product recommendations.
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Warehouse-USAL/wh-suggestions/internal/api"
	"github.com/Warehouse-USAL/wh-suggestions/internal/config"
)

func main() {
	cfg := config.Load()
	// TODO(RFC §5.2): start the Kafka consumer (topic cfg.KafkaTopic, group cfg.KafkaGroupID).
	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           api.NewRouter(),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("wh-suggestions listening on %s", cfg.HTTPAddr)
	log.Fatal(srv.ListenAndServe())
}
