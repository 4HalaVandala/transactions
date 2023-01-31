package main

import (
	"github.com/4halavandala/transactions/consumer/pkg/consts"
	"github.com/4halavandala/transactions/consumer/pkg/handler"
	"github.com/4halavandala/transactions/consumer/pkg/utils"
)

func main() {
	connectionString := utils.GetEnvVar("RMQ_URL")

	exampleQueue := utils.RMQConsumer{
		Queue:            consts.EXAMPLE_QUEUE,
		ConnectionString: connectionString,
		MsgHandler:       handler.HandleExample,
	}
	// Start consuming message on the specified queues
	forever := make(chan bool)

	go exampleQueue.Consume()

	// Multiple listeners can be specified here

	<-forever
}
