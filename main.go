package main

import (
	"BWA/auth"
	"BWA/campaign"
	"BWA/handler"
	"BWA/helper"
	"BWA/rpcp"
	"BWA/transactions"
	"BWA/user"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	godotenv.Load(`.env`)
	dsn := os.Getenv(`MYSQL_DSN`)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db1, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err.Error())

	}

	repoSQL := user.NewRepositorySQL(db1)

	//userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	TransactionRepo := transactions.NewRepository(db)

	userService := user.NewService(repoSQL)
	authService := auth.NewService()
	campaignService := campaign.NewService(campaignRepository)
	transactionService := transactions.NewService(TransactionRepo)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionsHandler := handler.NewTransaction(transactionService)

	go func() {
		const port = `:9090`
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost"+port))
		if err != nil {
			log.Fatal(`failed to listen on port ` + port)
		}
		grpcServer := grpc.NewServer()
		rpcp.RegisterUserServiceServer(grpcServer, userHandler)
		log.Fatal(grpcServer.Serve(lis))
	}()

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleWare(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaign", authMiddleWare(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaign/:id", authMiddleWare(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleWare(authService, userService), campaignHandler.UploadImage)

	api.GET("/campaign/:id/transactions", authMiddleWare(authService, userService), transactionsHandler.GetTransactions)

	router.Run()

}

func authMiddleWare(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var tokenString string
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
