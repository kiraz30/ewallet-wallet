package cmd

import (
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/api"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/repository"
	"ewallet-wallet/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	dependency := dependencyInject()
	healthCheckSVC := &services.HealthCheck{}
	healtCheckAPI := &api.HealthCheck{
		HealthCheckServices: healthCheckSVC,
	}

	r := gin.Default()
	r.GET("/health", healtCheckAPI.HealthChecHandlerHTTP)

	walletV1 := r.Group("/wallet/v1")
	walletV1.POST("/", dependency.WalletApi.Create)

	err := r.Run(":" + helpers.GetEnv("PORT", "8081"))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	HealtyCheckApi   interfaces.IHealthCheckApi
	WalletRepository interfaces.IWalletRepossitory
	WalletApi        interfaces.IWalletApi
}

func dependencyInject() Dependency {
	healtyCheckSVC := &services.HealthCheck{}
	healtyCheckAPI := &api.HealthCheck{
		HealthCheckServices: healtyCheckSVC,
	}

	walletRepository := &repository.WalletRepository{
		DB: helpers.DB,
	}
	walletSVC := &services.WalletService{
		WalletRepository: walletRepository,
	}
	walletAPI := &api.WalletApi{
		WalletService: walletSVC,
	}

	return Dependency{
		HealtyCheckApi:   healtyCheckAPI,
		WalletRepository: walletRepository,
		WalletApi:        walletAPI,
	}

}
