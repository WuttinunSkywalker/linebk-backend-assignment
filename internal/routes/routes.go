package routes

import (
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/account"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/auth"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/banner"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/debit"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/transaction"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/api/user"
	"github.com/WuttinunSkywalker/linebk-backend-assignment/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	authHandler auth.AuthHandler,
	userHandler user.UserHandler,
	bannerHandler banner.BannerHandler,
	transactionHandler transaction.TransactionHandler,
	debitHandler debit.DebitHandler,
	accountHandler account.AccountHandler,
	authMiddleware middleware.AuthMiddleware,
) {
	api := r.Group("/api")

	// Public routes
	authRoutes := api.Group("/auth")
	{
		authRoutes.POST("/login", authHandler.LoginWithPin)
	}

	publicUserRoutes := api.Group("/users")
	{
		publicUserRoutes.GET("/:userid/preview", userHandler.GetUserPreview)
	}

	// Protected routes
	protected := api.Group("")
	protected.Use(authMiddleware.ValidateToken())

	userRoutes := protected.Group("/users")
	{
		userRoutes.GET("/me", userHandler.GetMe)
		userRoutes.GET("/me/greetings", userHandler.GetMyGreeting)
	}

	bannerRoutes := protected.Group("/banners")
	{
		bannerRoutes.GET("", bannerHandler.GetMyBanners)
	}

	transactionRoutes := protected.Group("/transactions")
	{
		transactionRoutes.GET("", transactionHandler.GetMyTransactions)
	}

	debitRoutes := protected.Group("/debits")
	{
		debitRoutes.GET("", debitHandler.GetMyDebitCards)
	}

	accountRoutes := protected.Group("/accounts")
	{
		accountRoutes.GET("", accountHandler.GetMyAccounts)
	}

	// @Summary Health check
	// @Description Check if the API is running
	// @Tags health
	// @Accept json
	// @Produce json
	// @Success 200 {object} map[string]string "API is healthy"
	// @Router /health [get]
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
