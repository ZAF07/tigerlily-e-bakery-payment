package router

import (
	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rest/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {
	
	// Set CORS config
	r.Use(cors.New(cors.Config{
		AllowCredentials: false,
		// AllowAllOrigins: true,
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTION", "HEAD", "PATCH", "COMMON"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Authorization", "accept", "origin", "Referer", "User-Agent"},
	}))

	// r.Use(middleware.CORSMiddleware())

	// Checkout API Endpoint
	checkoutAPI := controller.NewCheckoutAPI()
	{
		checkOut := r.Group("checkout")
		{
			checkOut.POST("", checkoutAPI.Checkout)
		}
	}

	return r
}