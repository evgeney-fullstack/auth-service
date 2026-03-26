package handler

import (
	"github.com/evgeney-fullstack/auth-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Handler handles HTTP requests and manages routing.
// Contains dependencies required for request handlers (future fields).
type Handler struct {
	services     *service.Service // Service layer dependency for business logic
	errorHandler *ErrorHandler    // Error handler with logging capabilities
}

// NewHandler creates and returns a new Handler instance.
// Constructor function for initializing a handler with possible dependencies.
//
// Parameters:
//   - services: reference to the service layer that provides business logic
//   - logger: logrus logger for error logging and debugging
//
// Returns:
//   - *Handler: pointer to the newly created Handler instance
func NewHandler(services *service.Service, logger *logrus.Logger) *Handler {
	return &Handler{
		services:     services,
		errorHandler: NewErrorHandler(logger), // Initialize error handler with provided logger
	}
}

// InitRoutes configures and returns the Gin router with defined endpoints.
// Adds middleware and registers handlers for all API paths.
//
// Receiver:
//   - h: pointer to Handler instance containing required dependencies
//
// Returns:
//   - *gin.Engine: configured Gin router with all registered routes
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New() // Create a new Gin router instance without default middleware

	// Create a route group for auth-service endpoints with the prefix "/auth"
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)               // User registration endpoint
		auth.POST("/sign-in", h.signIn)               // User authentication endpoint
		auth.POST("/refresh-token", h.refreshHandler) // Token refresh endpoint
		auth.POST("/logout", h.logout)                // Single device logout endpoint
		auth.POST("/logout-all", h.logoutAll)         // Global logout (all devices) endpoint
		auth.GET("/validate", h.userIdentity)         // User identity validation endpoint
	}

	return router
}
