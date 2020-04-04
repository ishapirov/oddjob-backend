package handler

import (
	"OddJob/platform/request"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AvailableRequestGet(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestid := c.DefaultQuery("requestID", "0")
		var request request.Request
		query := "SELECT request_id,title,post,price,poster_id FROM requests WHERE request_id=?"

		err := db.QueryRow(query, requestid).Scan(&request.RequestID, &request.Title, &request.Post, &request.Price, &request.UserID)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		c.JSON(http.StatusOK, request)
	}
}