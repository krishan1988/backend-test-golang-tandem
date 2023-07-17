// Package handler contains handler functions which are related backend-test-golang server.
package handlers

import (
	"fmt"
	"net/http"

	"github.com/KryptoKnight/backend-test-golang/api/server/schema"
	"github.com/KryptoKnight/backend-test-golang/models"
	"github.com/KryptoKnight/backend-test-golang/service"
	"github.com/gin-gonic/gin"
)

// GetRetrieveFactsHandler get retrieve handler returns facts by based on the given filters.
func GetRetrieveFactsHandler(factService service.FactService) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		reqCtx := ctx.Request.Context()

		filter, err := models.NewFactFilterBuilder().
			SetTypes(ctx.QueryArray("type")).
			SetFound(ctx.Query("found")).
			SetLimit(ctx.Query("limit")).
			SetPage(ctx.Query("page")).
			Get()

		if err != nil {
			errSchema := schema.Error{
				Message: fmt.Sprintf("errors: %s", err),
			}
			ctx.JSON(http.StatusBadRequest, errSchema)
			return
		}
		facts, err := factService.Retrieve(reqCtx, filter)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		resp := schema.RetrieveFactsResponse{Facts: facts}
		ctx.JSON(http.StatusOK, resp)
	}
}
