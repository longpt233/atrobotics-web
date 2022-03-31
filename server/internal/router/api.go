package router

import (
	"atro/internal/handler"
	"atro/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

//RunAPI ->route setup
func RunAPI(address string) error {

	r := gin.Default()

	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true

	// r.Use(cors.New(config))
	// r.Use(corsMiddleware())

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Our Mini Ecommerce")
	})

	apiRoutes := r.Group("/api/v1")

	productHandler := handler.NewProductHandler()
	productCategoryHandler := handler.NewProductCategoryHandler()
	userHandler := handler.NewUserHandler()
	orderHandler := handler.NewOrderHandler()

	// api cho user
	userRoutes := apiRoutes.Group("/user")
	{
		// unauthorize api

		// đăng nhập đăng kí
		userRoutes.POST("/login", userHandler.SignInUser)
		userRoutes.POST("/register", userHandler.AddUser)
		userRoutes.POST("/logout", nil)

		// xem liên quan
		userRoutes.GET("/products/", productHandler.GetAllProduct)
		userRoutes.GET("/products/:id", productHandler.GetProduct)
		userRoutes.GET("/categories/", productCategoryHandler.GetAllProductCategories)
		userRoutes.GET("/categories/:id", productCategoryHandler.GetProductCategory)

		// authorize api
		userAuth := userRoutes.Group("/auth", middleware.AuthorizeJWT())
		userAuth.GET("/info", userHandler.GetUser)
		userAuth.PUT("/info", userHandler.UpdateUser)
		userAuth.POST("/change-password", userHandler.ChangePassword)

		// create order . chỉ cho tạo
		userAuth.POST("/orders", orderHandler.OrderProduct) // gửi lên cái là chốt đơn.

	}

	// api cho admin
	adminRouter := apiRoutes.Group("/admin")
	{
		// unauthorize api

		// authorize api
		// adminAuth := adminRouter.Group("/auth", middleware.AuthorizeJWT(), middleware.IsAdmin())
		adminAuth := adminRouter.Group("/auth")

		// category
		adminAuth.POST("/categories/", productCategoryHandler.AddProductCategory)
		adminAuth.DELETE("/categories/:id", productCategoryHandler.DeleteProductCategory)
		adminAuth.PUT("/categories/:id", productCategoryHandler.UpdateProductCategory)

		// product
		adminAuth.POST("/products/", productHandler.AddProduct)
		adminAuth.DELETE("/products/:id", productHandler.DeleteProduct)
		adminAuth.PUT("/products/:id", productHandler.UpdateProduct)

		// order info. không cho del, k cho tạo
		adminAuth.GET("/orders", orderHandler.GetAllOrderProduct)
		adminAuth.GET("/orders/:id", orderHandler.GetOrderProduct)
		adminAuth.PUT("/orders/:id", orderHandler.UpdateOrderProduct)

		// upload file
		adminAuth.POST("/file-uploads/single-file", handler.SingleFile)

	}
	return r.Run(address)

}

// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		fmt.Println("via midddleware")

// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(200)
// 		}

// 		// c.Next()

// 	}
// }
