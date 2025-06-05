package app

import "github.com/gin-gonic/gin"

func StartApi() {
	// pakai gin untuk api backend

	r := gin.Default()

	// panggil router
	Router(r)

	r.Run(":8080")

}
