package main

import (
	"bestHabit/component"
	"bestHabit/middleware"
	"bestHabit/modules/user/usertransport/ginuser"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func ConnectToDB(dns string) *sqlx.DB {
	db, err := sqlx.Connect("mysql", dns)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	// Check connection to DB by Ping
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}
	fmt.Println("Connected to the database!")

	return db
}

func runServer(db *sqlx.DB, secretKey string) {

	appCtx := component.NewAppContext(db, secretKey)
	fmt.Print(appCtx)

	router := gin.Default()

	router.Use(middleware.Recover(appCtx))

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ping OK!",
		})
	})

	log_and_register := router.Group("/")
	{
		log_and_register.POST("/register", ginuser.BasicRegister(appCtx))
	}

	router.Run(":8080")
}

func main() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	dns := os.Getenv("DB_CONNECTION_STR")
	secretKet := os.Getenv("SECRET_KEY")

	db := ConnectToDB(dns)
	runServer(db, secretKet)
}
