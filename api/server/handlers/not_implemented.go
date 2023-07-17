// Package handler contains handler functions which are related backend-test-golang server.
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNotImplementedHandler returns not implemented handler.
func GetNotImplementedHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusNotImplemented)
	}
}
