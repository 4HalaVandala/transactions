package handler

import (
	"net/http"

	"github.com/4halavandala/transactions/producer/domain/model"
	"github.com/4halavandala/transactions/producer/pkg/consts"
	"github.com/4halavandala/transactions/producer/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SendMessageInQ(c *gin.Context) {
	var msg model.Message

	request_id := c.GetString("x-request-id")

	// Bind request payload with our model
	if binderr := c.ShouldBindJSON(&msg); binderr != nil {

		log.Error().Err(binderr).Str("request_id", request_id).
			Msg("Error occurred while binding request data")

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": binderr.Error(),
		})
		return
	}

	connectionString := utils.GetEnvVar("RMQ_URL")

	rmqProducer := utils.RMQProducer{
		Queue:            consts.EXAMPLE_QUEUE,
		ConnectionString: connectionString,
	}

	rmqProducer.PublishMessage("text/plain", []byte(msg.Message))

	c.JSON(http.StatusOK, gin.H{
		"response": "Message received",
	})

}
