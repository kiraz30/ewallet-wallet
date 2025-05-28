package cmd

import (
	"ewallet-wallet/external"
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/api"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/repository"
	"ewallet-wallet/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	d := dependencyInject()
	healthCheckSVC := &services.HealthCheck{}
	healtCheckAPI := &api.HealthCheck{
		HealthCheckServices: healthCheckSVC,
	}

	r := gin.Default()
	r.GET("/health", healtCheckAPI.HealthChecHandlerHTTP)

	walletV1 := r.Group("/wallet/v1")
	walletV1.POST("/", d.WalletApi.Create)
	walletV1.PUT("/balance/credit", d.MiddlewareValidateToken, d.WalletApi.CreaditBalance)
	walletV1.PUT("/balance/debit", d.MiddlewareValidateToken, d.WalletApi.DebitBalance)
	walletV1.GET("/balance", d.MiddlewareValidateToken, d.WalletApi.GetWalletBalance)
	err := r.Run(":" + helpers.GetEnv("PORT", "8081"))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	HealtyCheckApi   interfaces.IHealthCheckApi
	WalletRepository interfaces.IWalletRepossitory
	WalletApi        interfaces.IWalletApi
	External         interfaces.IExternal
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
	external := &external.External{}

	return Dependency{
		HealtyCheckApi:   healtyCheckAPI,
		WalletRepository: walletRepository,
		WalletApi:        walletAPI,
		External:         external,
	}

}
