package app

import (
	"log"
	"os"
	"strconv"
	"tes1/varglobal"

	"github.com/joho/godotenv"
)

func Loadconfig() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logFile, err := os.Create("app.log")
	if err != nil {
		log.Fatal("failed to create log file")
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("App started")

	// load environment variable form file and set to file

	// retrieve and print the environment variable
	// Akses variabel environment
	varglobal.DatabaseName = os.Getenv("DATABASE_NAME")
	varglobal.DatabaseUser = os.Getenv("DATABASE_USER")
	varglobal.DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	varglobal.DatabasePort = os.Getenv("DATABASE_PORT")
	varglobal.Mainport, _ = strconv.Atoi(os.Getenv("MAIN_PORT"))
	varglobal.DatabaseHost = os.Getenv("DATABASE_HOST")

	log.Println("Database Name:", varglobal.DatabaseName)
	log.Println("Database User:", varglobal.DatabaseUser)
	log.Println("Database Password:", varglobal.DatabasePassword)
	log.Println("Database Port:", varglobal.DatabasePort)
	log.Println("Main Port:", varglobal.Mainport)
	log.Println("Database Host:", varglobal.DatabaseHost)

}
