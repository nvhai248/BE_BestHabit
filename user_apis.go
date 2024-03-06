package main

import (
	"bestHabit/component"
	proto "bestHabit/generatedProto/proto/userservice"
	"bestHabit/middleware"
	"bestHabit/modules/upload/uploadtransport/ginupload"
	"bestHabit/modules/user/userstorage/usergrpcserver"
	"bestHabit/modules/user/usertransport/ginuser"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

func ApisUser(appCtx component.AppContext, ver *gin.RouterGroup, db *sqlx.DB) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	user := ver.Group("/users", middleware.RequireAuth(appCtx), middleware.IsVerifiedUser(appCtx))
	{
		user.PATCH("/profile", ginuser.UpdateProfile(appCtx, cc))
		user.GET("/profile", ginuser.GetProfile(appCtx))
		user.POST("/upload", ginupload.Upload(appCtx))
		user.PATCH("/change-password", ginuser.ChangePassword(appCtx, cc))
		user.PATCH("/reset-password", ginuser.ResetPassword(appCtx))
		user.PATCH("/device-token", ginuser.UpdateDeviceToken(appCtx))
	}

	verifyUser := ver.Group("/users", middleware.RequireAuth(appCtx))
	{
		verifyUser.POST("/send-verification", ginuser.SendVerification(appCtx))
		verifyUser.PATCH("/verify/:token", middleware.CompareIdBeforeVerify(appCtx), ginuser.Verify(appCtx))
	}

	// Create a listener on TCP port
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	proto.RegisterUserServiceServer(s, usergrpcserver.NewGRPCServer(db))

	go func() {
		// Serve gRPC Server
		log.Println("Serving gRPC on ", address)
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
}
