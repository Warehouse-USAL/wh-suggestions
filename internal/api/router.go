// Package api exposes the HTTP surface of the service (RFC §5.4).
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewRouter wires every route. The recommendation endpoints are scaffolded and
// answer 501 until the service is implemented.
func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/healthz", healthz)

	// TODO(RFC §5.4): JWT middleware (HS256, shared secret with wh-backend).
	v1 := r.Group("/v1/recommendations")
	v1.POST("/cart", notImplemented)
	v1.GET("/user/:userId", notImplemented)
	v1.GET("/hot", notImplemented)

	return r
}

// TODO(RFC §5.4): check that Redis and the Kafka consumer are alive.
func healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func notImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
