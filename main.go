package main

import (
	"bestHabit/component"
	"bestHabit/component/mailprovider"
	"bestHabit/component/uploadprovider"
	"bestHabit/docs"
	"bestHabit/middleware"
	"bestHabit/modules/challenge/challengetransport/ginchallenge"
	"bestHabit/modules/habit/habittransport/ginhabit"
	"bestHabit/modules/participant/participanttransport/ginparticipant"
	"bestHabit/modules/task/tasktransport/gintask"
	"bestHabit/modules/upload/uploadtransport/ginupload"
	"bestHabit/modules/user/usertransport/ginuser"
	"bestHabit/pubsub/pblocal"
	"bestHabit/skio"
	"bestHabit/subscriber"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

func runServer(db *sqlx.DB, secretKey string, s3upProvider uploadprovider.UploadProvider, gmailSender mailprovider.EmailSender) {

	appCtx := component.NewAppContext(db, secretKey, s3upProvider, pblocal.NewPubSub(), gmailSender)

	router := gin.Default()

	router.Use(middleware.Recover(appCtx))

	// start subscriber
	if err := subscriber.NewEngine(appCtx).Start(); err != nil {
		log.Fatal(err)
	}

	rtEngine := skio.NewEngine()

	if err := rtEngine.Run(appCtx, router); err != nil {
		log.Fatalln("Failed to start server: ", err)
	}

	docs.SwaggerInfo.BasePath = "/api"

	routerAPIS := router.Group("/api")

	log_and_register := routerAPIS.Group("/")
	{
		log_and_register.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Ping OK!!!!",
			})
		})
		log_and_register.POST("/register", ginuser.BasicRegister(appCtx))
		log_and_register.POST("/login", ginuser.BasicLogin(appCtx))
	}

	user := routerAPIS.Group("/users", middleware.RequireAuth(appCtx))
	{
		user.PATCH("/profile", ginuser.UpdateProfile(appCtx))
		user.GET("/profile", ginuser.GetProfile(appCtx))
		user.POST("/upload", ginupload.Upload(appCtx))
		user.POST("/send-verification", ginuser.SendVerification(appCtx))
		user.PATCH("/verify/:token", middleware.CompareIdBeforeVerify(appCtx), ginuser.Verify(appCtx))
	}

	task := routerAPIS.Group("/tasks", middleware.RequireAuth(appCtx))
	{
		task.POST("/", gintask.CreateTask(appCtx))
		task.GET("/", gintask.ListTaskByConditions(appCtx))
		task.GET("/:id", gintask.FindTask(appCtx))
		task.PATCH("/:id", gintask.UpdateTask(appCtx))
		task.DELETE("/:id", gintask.SoftDeleteTask(appCtx))
	}

	habit := routerAPIS.Group("/habits", middleware.RequireAuth(appCtx))
	{
		habit.POST("/", ginhabit.CreateHabit(appCtx))
		habit.GET("/", ginhabit.ListHabitByConditions(appCtx))
		habit.GET("/:id", ginhabit.FindHabit(appCtx))
		habit.PATCH("/:id", ginhabit.UpdateHabit(appCtx))
		habit.DELETE("/:id", ginhabit.SoftDeleteHabit(appCtx))
		habit.PATCH("/:id/confirm-completed", ginhabit.AddCompletedDate(appCtx))
	}

	challenge := routerAPIS.Group("/challenges", middleware.RequireAuth(appCtx))
	{
		challenge.GET("/", ginchallenge.ListChallengeByConditions(appCtx))
		challenge.GET("/:id", ginchallenge.FindChallenge(appCtx))

		challenge.POST("/:id/user-join", ginparticipant.CreateParticipant(appCtx))
		challenge.DELETE("/:id/user-cancel", ginparticipant.CancelParticipant(appCtx))
		challenge.GET("/participants", ginparticipant.ListChallengeJoined(appCtx))
		challenge.PATCH("/participants/:id", ginparticipant.UpdateParticipant(appCtx))
		challenge.GET("/participants/:id", ginparticipant.FindParticipant(appCtx))
	}

	challengeAdmin := routerAPIS.Group("/challenges", middleware.RequireAuth(appCtx), middleware.RequireRoles(appCtx, "admin"))
	{
		challengeAdmin.POST("/", ginchallenge.CreateChallenge(appCtx))
		challengeAdmin.PATCH("/:id", ginchallenge.UpdateChallenge(appCtx))
		challengeAdmin.DELETE("/:id", ginchallenge.DeleteChallenge(appCtx))
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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

	// get Sender
	appName := os.Getenv("SENDER_APP_NAME")
	appEmailPw := os.Getenv("SENDER_APP_PW")
	appEmailAdd := os.Getenv("SENDER_EMAIL_ADDRESS")
	gmailSender := mailprovider.NewGmailSender(appName, appEmailAdd, appEmailPw)
	runServer(db, secretKet, s3upProvider, gmailSender)
}
