package handler

import (
	"event-driven/models"
	"event-driven/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlePostMessage handles the POST request to send messages
func HandlePostMessage(c *gin.Context) {
	var msg models.Message

	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Both 'user' and 'message' fields are required and must be valid JSON.",
		})
		return
	}

	// Initialize RabbitMQ connection
	conn, ch, err := utils.InitRabbitMQ()
	if err != nil {
		log.Printf("Error initializing RabbitMQ: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error initializing RabbitMQ.",
		})
		return
	}
	defer conn.Close()
	defer ch.Close()

	err = utils.PublishMessage(ch, "PublishQueue", msg)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to publish message to the Publish service.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Message successfully sent to Publish service.",
	})
}
