package main

import (
	"tes1/app"
	"tes1/dbku"
)

func main() {
	app.Loadconfig() // Load configuration from .env file
	dbku.InitDB()
	app.StartApi()

}
