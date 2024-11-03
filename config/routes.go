package config

import (
	"github.com/darkphotonKN/ecommerce-server-go/internal/product"
	"github.com/darkphotonKN/ecommerce-server-go/internal/user"
	"github.com/gin-gonic/gin"
)

/**
* Sets up API prefix route and all routers.
**/
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// base route
	api := router.Group("/api")

	// -- USER --

	// --- User Setup ---
	userRepo := user.NewUserRepository(DB)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// --- User Routes ---
	userRoutes := api.Group("/user")
	userRoutes.GET("/:id", userHandler.GetUserByIdHandler)
	userRoutes.POST("/", userHandler.CreateUserHandler)

	// -- PRODUCT --

	// --- Product Setup ---
	productRepo := product.NewProductRepository(DB)
	productService := product.NewProductService(productRepo)
	productHandler := product.NewProductHandler(productService)

	// --- Product Routes ---
	productRoutes := api.Group("/product")
	productRoutes.GET("/", productHandler.GetProductsHandler)
	productRoutes.GET("/trending", productHandler.GetTrendingProductsHandler)
	productRoutes.POST("/", productHandler.CreateProductsHandler)
	return router
}
