package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/4halavandala/transactions/tx-app/pkg/consts"
	"github.com/4halavandala/transactions/tx-app/pkg/handler"
	"github.com/4halavandala/transactions/tx-app/pkg/repository"
	"github.com/4halavandala/transactions/tx-app/pkg/service"
	"github.com/4halavandala/transactions/tx-app/pkg/utils"
	"github.com/4halavandala/transactions/tx-app/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error while lodaing .env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to init database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(transactions.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != http.ErrServerClosed {
			logrus.Fatalf("error occured while running server: %s", err.Error())
		}
	}()
	logrus.Print("transaction app Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("transaction app Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

	// Rabbit mq connection
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

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("Config")

	return viper.ReadInConfig()
}
