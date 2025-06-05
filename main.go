package main

import (
	"tes1/app"
	"tes1/dbku"
)

func main() {
	dbku.InitDB()
	app.StartApi()

}
