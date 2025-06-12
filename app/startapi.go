package app

import (
	"strconv"
	"tes1/varglobal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartApi() {
	// pakai gin untuk api backend

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "http://localhost:8080", "http://localhost"}, // Ganti dengan origin tertentu jika perlu
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// panggil router
	Router(r)
	port := strconv.Itoa(varglobal.Mainport)

	r.Run(":" + port) // listen and serve on

}
