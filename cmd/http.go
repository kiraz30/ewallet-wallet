package cmd

import (
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/api"
	"ewallet-wallet/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	healthCheckSVC := &services.HealthCheck{}
	healtCheckAPI := &api.HealthCheck{
		HealthCheckServices: healthCheckSVC,
	}

	r := gin.Default()
	r.GET("/health", healtCheckAPI.HealthChecHandlerHTTP)
	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}
