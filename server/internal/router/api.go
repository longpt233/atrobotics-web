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

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Our Mini Ecommerce")
	})

	apiRoutes := r.Group("/api/v1")

	productHandler := handler.NewProductHandler()
	productCategoryHandler := handler.NewProductCategoryHandler()
	userHandler := handler.NewUserHandler()

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
		userAuth := userRoutes.Group("/auth",middleware.AuthorizeJWT()) 
		userAuth.GET("", userHandler.GetUser) 
	}


	// api cho admin
	adminRouter := apiRoutes.Group("/admin")
	{
		// unauthorize api 

		// authorize api
		adminAuth := adminRouter.Group("/auth", middleware.AuthorizeJWT(), middleware.IsAdmin())

		// category
		adminAuth.POST("/categories/", productCategoryHandler.AddProductCategory)
		adminAuth.DELETE("/categories/:id", productCategoryHandler.DeleteProductCategory)
		adminAuth.PUT("/categories/:id", productCategoryHandler.UpdateProductCategory)

		// product
		adminAuth.POST("/products/", productHandler.AddProduct)
		adminAuth.DELETE("/products/:id", productHandler.DeleteProduct)
		adminAuth.PUT("/products/:id",productHandler.UpdateProduct)
		
		// order info

		// upload file
		adminAuth.POST("/file-uploads/single-file", handler.SingleFile)



	}

	return r.Run(address)

}
