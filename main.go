package main

import (
	"bestHabit/component"
	"bestHabit/component/uploadprovider"
	"bestHabit/middleware"
	"bestHabit/modules/habit/habiztransport/ginhabit"
	"bestHabit/modules/task/tasktransport/gintask"
	"bestHabit/modules/upload/uploadtransport/ginupload"
	"bestHabit/modules/user/usertransport/ginuser"
	"bestHabit/pubsub/pblocal"
	"bestHabit/subscriber"
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

func runServer(db *sqlx.DB, secretKey string, s3upProvider uploadprovider.UploadProvider) {

	appCtx := component.NewAppContext(db, secretKey, s3upProvider, pblocal.NewPubSub())

	router := gin.Default()

	router.Use(middleware.Recover(appCtx))

	// start subscriber
	if err := subscriber.NewEngine(appCtx).Start(); err != nil {
		log.Fatal(err)
	}

	log_and_register := router.Group("/")
	{
		log_and_register.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Ping OK!",
			})
		})
		log_and_register.POST("/register", ginuser.BasicRegister(appCtx))
		log_and_register.POST("/login", ginuser.BasicLogin(appCtx))
	}

	user := router.Group("/users", middleware.RequireAuth(appCtx))
	{
		user.PATCH("/profile", ginuser.UpdateProfile(appCtx))
		user.GET("/profile", ginuser.GetProfile(appCtx))
		user.POST("/upload", ginupload.Upload(appCtx))
	}

	task := router.Group("/tasks", middleware.RequireAuth(appCtx))
	{
		task.POST("/", gintask.CreateTask(appCtx))
		task.GET("/", gintask.ListTaskByConditions(appCtx))
		task.GET("/:id", gintask.FindTask(appCtx))
		task.PATCH("/:id", gintask.UpdateTask(appCtx))
	}

	habit := router.Group("/habits", middleware.RequireAuth(appCtx))
	{
		habit.POST("/", ginhabit.CreateHabit(appCtx))
		habit.GET("/", ginhabit.ListHabitByConditions(appCtx))
		habit.GET("/:id", ginhabit.FindHabit(appCtx))
		habit.PATCH("/:id", ginhabit.UpdateTask(appCtx))
		habit.DELETE("/:id", ginhabit.SoftDeleteTask(appCtx))
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
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3ApiKey := os.Getenv("S3ApiKey")
	s3Secret := os.Getenv("S3Secret")
	s3Domain := os.Getenv("S3Domain")
	s3upProvider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3ApiKey, s3Secret, s3Domain)

	db := ConnectToDB(dns)
	runServer(db, secretKet, s3upProvider)
}
