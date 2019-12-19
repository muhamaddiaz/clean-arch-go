package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"log"
	"runtime"
	delivery "service-account/account/delivery/http"
	"service-account/account/repository"
	"service-account/account/usecase"
)

func init() {
	runtime.GOMAXPROCS(2)

	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigFile("app.config.json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
}

func catchError() {
	if e := recover(); e != nil {
		log.Fatal("Error Occured!, ", e)
	}
}

func main() {
	defer catchError()

	e := echo.New()

	database := viper.GetStringMapString("database")
	host := database["host"]
	port := database["port"]
	user := database["user"]
	dbname := database["dbname"]
	password := database["password"]
	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err := gorm.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	ar := repository.NewPostgresAccountRepository(db)
	au := usecase.NewAccountUsecase(ar)
	delivery.NewAccountRouteHandler(e, au)

	log.Fatal(e.Start(":9090"))
}