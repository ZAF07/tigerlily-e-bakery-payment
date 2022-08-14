package env

import (
	"fmt"
	"os"

	"github.com/Tiger-Coders/tigerlily-payment/internal/pkg/logger"
)

func SetEnv() {

	logs := logger.NewLogger()

	logs.InfoLogger.Println("Setting ENV ... ")
	serverENV := os.Args[1]
	dbHost := os.Args[2]
	dbUser := os.Args[3]
	dbPassword := os.Args[4]
	dbName := os.Args[5]
	dbSSL := os.Args[6]
	dbPort := os.Args[7]

	// port := strconv.Itoa(dbPort)
	os.Setenv("serverenv", serverENV)
	os.Setenv("dbHost", dbHost)
	os.Setenv("dbUser", dbUser)
	os.Setenv("dbPassword", dbPassword)
	os.Setenv("dbName", dbName)
	os.Setenv("dbPort", dbPort)
	os.Setenv("dbSSL", dbSSL)
	logs.InfoLogger.Println("DONE SETTING ENV")
	// fmt.Println("DONE SETTING ENV")
}

func GetDBEnv() (dbString string) {
	logs := logger.NewLogger()

	env := os.Getenv("serverenv")
	host := os.Getenv("dbHost")
	user := os.Getenv("dbUser")
	password := os.Getenv("dbPassword")
	dbname := os.Getenv("dbName")
	port := os.Getenv("dbPort")
	sslMode := os.Getenv("dbSSL")

	switch env {
	case "PROD":
		logs.InfoLogger.Println("RUNNING ON PRODUCTION MODE")

		dbString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslMode)
		return
	default:
		logs.InfoLogger.Println("RUNNING ON DEVELOPMENT MODE")
		dbString = fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s", host, user, dbname, port, sslMode)
	}
	return
}
