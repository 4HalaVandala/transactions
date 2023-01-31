package app

import (
	"github.com/4halavandala/transactions/producer/pkg/handler"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Function to setup the app object
func SetupApp() *gin.Engine {
	log.Info().Msg("Initializing service")

	// Create barebone engine
	app := gin.New()
	// Add default recovery middleware
	app.Use(gin.Recovery())

	// disabling the trusted proxy feature
	app.SetTrustedProxies(nil)

	// Setup routers
	log.Info().Msg("Setting up routers")
	handler.SetupRouters(app)

	return app
}
