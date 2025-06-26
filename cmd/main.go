package main

import (
	"fmt"

	_ "github.com/WuttinunSkywalker/linebk-backend-assignment/docs"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/account"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/auth"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/banner"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/debit"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/transaction"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/user"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/config"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/database"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/middleware"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/routes"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title LineBK Backend Assignment API
// @version 1.0
// @description This is the API documentation for LineBK Backend Assignment
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token example: "Bearer {token}.
func main() {
	cfg := config.GetConfig()

	logger.Init(cfg.Log)

	db := database.New(cfg.DB)
	defer db.Close()

	// repositories
	userRepo := user.NewUserRepository(db)
	bannerRepo := banner.NewBannerRepository(db)
	transactionRepo := transaction.NewTransactionRepository(db)
	debitRepo := debit.NewDebitRepository(db)
	accountRepo := account.NewAccountRepository(db)

	// usecases
	authUsecase := auth.NewAuthUsecase(userRepo, cfg.JWT)
	userUsecase := user.NewUserUsecase(userRepo)
	bannerUsecase := banner.NewBannerUsecase(bannerRepo)
	transactionUsecase := transaction.NewTransactionUsecase(transactionRepo)
	debitUsecase := debit.NewDebitUsecase(debitRepo)
	accountUsecase := account.NewAccountUsecase(accountRepo)

	// handlers
	authHandler := auth.NewAuthHandler(authUsecase)
	userHandler := user.NewUserHandler(userUsecase)
	bannerHandler := banner.NewBannerHandler(bannerUsecase)
	transactionHandler := transaction.NewTransactionHandler(transactionUsecase)
	debitHandler := debit.NewDebitHandler(debitUsecase)
	accountHandler := account.NewAccountHandler(accountUsecase)

	// middleware
	authMiddleware := middleware.NewAuthMiddleware(cfg.JWT)

	r := gin.Default()

	r.Use(cors.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.ErrorHandler())

	// docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterRoutes(
		r,
		authHandler,
		userHandler,
		bannerHandler,
		transactionHandler,
		debitHandler,
		accountHandler,
		authMiddleware,
	)

	portStr := fmt.Sprintf("%d", cfg.Port)
	logger.Infof("Starting server on port %s", portStr)
	if err := r.Run(":" + portStr); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
