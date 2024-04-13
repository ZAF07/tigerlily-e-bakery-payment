package router

import (
	"github.com/Tiger-Coders/tigerlily-payment/api/rest/controller"
	"github.com/Tiger-Coders/tigerlily-payment/internal/injection"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {
	config := injection.GetGeneralConfig()
	// Set CORS config
	r.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowAllOrigins:  true,
		// AllowOrigins: []string{"http://localhost:8080"},
		// AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTION", "HEAD", "PATCH", "COMMON"},
		// AllowHeaders: []string{"Content-Type", "Content-Length", "Authorization", "accept", "origin", "Referer", "User-Agent"},
		AllowOrigins: config.ServerConfig.AllowOrigins,
		AllowMethods: config.ServerConfig.AllowMethods,
		AllowHeaders: config.ServerConfig.AllowHeaders,
	}))

	// r.Use(middleware.CORSMiddleware())

	// Checkout API Endpoint
	checkoutAPI := controller.NewCheckoutAPI()
	checkOut := r.Group("checkout")

	{
		checkOut.POST("/stripe-checkout-session", checkoutAPI.StripeCheckoutSession)
	}

	return r
}
