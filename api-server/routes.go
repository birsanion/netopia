package main

import (
	"github.com/Depado/ginprom"
	"github.com/birsanion/netopia/api-server/middlewares"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"gorm.io/gorm"
)

func RegisterHealthRoute(router *gin.Engine, db *gorm.DB, queue *RabbitMQConnection) {
	router.GET("/health", HealthCheckHandler(db, queue))
}

func RegisterPaymentsRoute(router *gin.Engine, db *gorm.DB, queue *RabbitMQConnection) {
	api := router.Group("")
	api.Use(middlewares.AuthenticationMiddleware())
	{
		api.POST("/payments", CreatePaymentHandler(db, queue))
	}
}

func RegisterRoutes(router *gin.Engine, queue *RabbitMQConnection, db *gorm.DB) {
	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		RequestHeaders: "Origin, Authorization, Content-Type, X-API-Key",
	}))

	p := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	router.Use(p.Instrument())

	RegisterHealthRoute(router, db, queue)
	RegisterPaymentsRoute(router, db, queue)
}
