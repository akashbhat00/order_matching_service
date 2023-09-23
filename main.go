package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"order_matching_service/controllers"
	"order_matching_service/db"
	"order_matching_service/repositories"
	"order_matching_service/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// CORSMiddleware ...
// CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}

func main() {
	//Start the default gin server
	r := gin.Default()
	r.Use(CORSMiddleware())

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	//Start PostgreSQL database
	//Example: db.GetDB() - More info in the models folder
	db.Init()

	gorm := db.DB.GetDB(db.DB{})

	sellerRepository := repositories.NewSellerRepository(gorm)
	sellerService := services.NewSellerService(sellerRepository)
	sellerController := controllers.NewSellerController(sellerService)
	buyerRepository := repositories.NewBuyerRepository(gorm)
	buyerService := services.NewBuyerService(buyerRepository)
	buyerController := controllers.NewBuyerController(buyerService)
	orderRepository := repositories.NewOrderRepository(gorm)
	orderService := services.NewOrderService(orderRepository)
	orderController := controllers.NewOrderController(orderService)
	v1 := r.Group("/v1")
	{
		/*** START USER ***/

		v1.POST("/sellers", sellerController.CreateSeller)
		v1.GET("/sellers/:seller_id", sellerController.GetSellerByID)
		v1.POST("/buyers", buyerController.CreateBuyer)
		v1.GET("/buyers/:buyer_id", buyerController.GetBuyerByID)
		v1.POST("/sellers/:seller_id/products", sellerController.AddProductToSellerCatalog)
		v1.POST("/orders", orderController.PlaceOrder)
		// v1.POST("/seller", sellerController.CreateSeller)
		// v1.POST("/seller", sellerController.CreateSeller)
		// v1.POST("/seller", sellerController.CreateSeller)
		// v1.POST("/seller", sellerController.CreateSeller)

		/*** START Article ***/
		// article := new(controllers.ArticleController)

		// v1.POST("/article", TokenAuthMiddleware(), article.Create)
		// v1.GET("/articles", TokenAuthMiddleware(), article.All)
		// v1.GET("/article/:id", TokenAuthMiddleware(), article.One)
		// v1.PUT("/article/:id", TokenAuthMiddleware(), article.Update)
		// v1.DELETE("/article/:id", TokenAuthMiddleware(), article.Delete)
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	port := os.Getenv("PORT")

	log.Printf("\n\n PORT: %s \n ENV: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("API_VERSION"))
	r.Run(":" + port)

}
