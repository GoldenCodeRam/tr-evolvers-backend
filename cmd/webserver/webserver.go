package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goldencoderam/tr-evolvers-backend/pkg/database"
)

func main() {
    conn := database.Connect()

    conn.AutoMigrate(&database.EnergyMeterBrand{})
    conn.AutoMigrate(&database.EnergyMeter{})
    conn.AutoMigrate(&database.EnergyMeterInstallation{})

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
