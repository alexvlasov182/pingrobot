package handlers

import (
	"net/http"

	"github.com/alexvlasov182/http/pingrobot/backend/backend/workerpool"
	"github.com/gin-gonic/gin"
)

func ResultsHandler(results chan workerpool.Result) gin.HandlerFunc {

	return func(c *gin.Context) {
		var resultsList []gin.H
		for result := range results {
			resultsList = append(resultsList, gin.H{
				"url":    result.URL,
				"status": result.StatusCode,
				"time":   result.ResponseTime.String(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"results": resultsList,
		})
	}
}
