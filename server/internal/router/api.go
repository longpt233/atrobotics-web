package router

import (
	"atro/internal/handler"
	"atro/internal/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// cors "github.com/rs/cors/wrapper/gin"
)

//RunAPI ->route setup
func RunAPI(address string) error {

	r := gin.Default()

	// cors config
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Our Mini Ecommerce")
	})

	apiRoutes := r.Group("/api/v1")

	productHandler := handler.NewProductHandler()
	productCategoryHandler := handler.NewProductCategoryHandler()
	userHandler := handler.NewUserHandler()
	orderHandler := handler.NewOrderHandler()
	bannerHandler := handler.NewBannerHandler()
	cartItems := handler.NewCartItemsHandler()

	// api cho user
	userRoutes := apiRoutes.Group("/user")
	// userRoutes.Use(cors.Default())
	{
		// unauthorize api

		// đăng nhập đăng kí
		userRoutes.POST("/login", userHandler.SignInUser)
		userRoutes.POST("/register", userHandler.AddUser)
		userRoutes.POST("/logout", nil)
		userRoutes.GET("/forgot-password",userHandler.ForgotPassword)

		// xem liên quan
		userRoutes.GET("/products", productHandler.GetAllProduct)
		userRoutes.GET("/products/:id", productHandler.GetProduct)
		userRoutes.GET("/products/by-category", productHandler.GetProductByCategory)
		userRoutes.GET("/categories", productCategoryHandler.GetAllProductCategories)
		userRoutes.GET("/categories/:id", productCategoryHandler.GetProductCategory)
		userRoutes.GET("/all-brand", productHandler.GetAllProductBrand)
		userRoutes.GET("/search",productHandler.SearchByShortDescription)

		// get banner
		userRoutes.GET("/banners/:id", bannerHandler.GetBanner)
		userRoutes.GET("/banners/top-3-newest", bannerHandler.GetTop3NewestBanner)

		// authorize api
		userAuth := userRoutes.Group("/auth", middleware.AuthorizeJWT())
		userAuth.GET("/info", userHandler.GetUser)
		userAuth.PUT("/info", userHandler.UpdateUser)
		userAuth.POST("/change-password", userHandler.ChangePassword)
		// create order . chỉ cho tạo
		userAuth.POST("/orders", orderHandler.OrderProduct) // gửi lên cái là chốt đơn.
		userAuth.GET("/cart/list", cartItems.GetCartItemsByUserId)
		userAuth.POST("/cart/add", cartItems.AddCartItems)
		userAuth.DELETE("/cart/:id", cartItems.DeleteCartItems)
		userAuth.PUT("/cart/:id",cartItems.UpdateCartItems)

	}

	// api cho admin
	adminRouter := apiRoutes.Group("/admin")
	{
		// unauthorize api

		// authorize api
		adminAuth := adminRouter.Group("/auth", middleware.AuthorizeJWT(), middleware.IsAdmin())

		//user 
		adminAuth.GET("/user/lists",userHandler.GetAllUser)
		// category
		adminAuth.POST("/categories", productCategoryHandler.AddProductCategory)
		adminAuth.DELETE("/categories/:id", productCategoryHandler.DeleteProductCategory)
		adminAuth.PUT("/categories/:id", productCategoryHandler.UpdateProductCategory)

		// product
		adminAuth.POST("/products", productHandler.AddProduct)
		adminAuth.DELETE("/products/:id", productHandler.DeleteProduct)
		adminAuth.PUT("/products/:id", productHandler.UpdateProduct)

		// order info. không cho del, k cho tạo
		adminAuth.GET("/orders", orderHandler.GetAllOrderProduct)
		adminAuth.GET("/orders/:id", orderHandler.GetOrderProduct)
		adminAuth.PUT("/orders/:id", orderHandler.UpdateOrderProduct)

		// upload file
		adminAuth.POST("/file-uploads/single-file", handler.SingleFile)

		//add banner
		adminAuth.POST("/banners", bannerHandler.AddBanner)
		adminAuth.PUT("/banners/:id", bannerHandler.UpdateBanner)
		adminAuth.DELETE("/banners/:id", bannerHandler.DeleteBanner)

	}
	return r.Run(address)

}
