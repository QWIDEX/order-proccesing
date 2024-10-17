package common

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func ReadJSON(c *gin.Context, data any) error {
	return json.NewDecoder(c.Request.Body).Decode(data)
}

func WriteError(c *gin.Context, err ApiError) {
	log.Println("Api error occurred: ", err)
	c.JSON(err.Code, gin.H{"error": err.Message})
}
