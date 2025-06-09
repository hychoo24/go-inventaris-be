package app

import (
	"strconv"
	"tes1/varglobal"

	"github.com/gin-gonic/gin"
)

func StartApi() {
	// pakai gin untuk api backend

	r := gin.Default()

	// panggil router
	Router(r)
	port := strconv.Itoa(varglobal.Mainport)

	r.Run(":" + port) // listen and serve on

}
