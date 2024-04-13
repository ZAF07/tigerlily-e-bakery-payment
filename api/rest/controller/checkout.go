package controller

import (
	"context"
	// "database/sql"
	"fmt"
	"net/http"

	"github.com/Tiger-Coders/tigerlily-payment/api/rpc"
	"github.com/Tiger-Coders/tigerlily-payment/internal/injection"
	"github.com/Tiger-Coders/tigerlily-payment/internal/pkg/logger"
	repo "github.com/Tiger-Coders/tigerlily-payment/internal/repository/checkout"
	"github.com/Tiger-Coders/tigerlily-payment/internal/service/checkout"
	"github.com/gin-gonic/gin"
)

type CheckoutAPI struct {
	// db   *sql.DB
	db   repo.CheckoutDBInterface
	logs logger.Logger
}

func NewCheckoutAPI() *CheckoutAPI {
	return &CheckoutAPI{
		// db:   db.NewDB(),
		db:   injection.GetPaymentDBInstance(),
		logs: *logger.NewLogger(),
	}
}

func (a CheckoutAPI) StripeCheckoutSession(c *gin.Context) {
	a.logs.InfoLogger.Println("[CONTROLLER] Checkout API running")

	var req rpc.CheckoutReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Printf("error binding req struct : %+v\n", err)
	}
	fmt.Printf("HERE FROM CONTROLLER : %+v\n", req.CheckoutItems[0])
	ctx := context.Background()
	service := checkout.NewCheckoutService(a.db)

	// PROPERLY HANDLE ERROR FOR WHEN DB ERROR
	resp, err := service.CustomCheckout(ctx, &req)
	if err != nil {
		a.logs.ErrorLogger.Println("[CONTROLLER] Error getting response")
		a.logs.InfoLogger.Printf("[CONTROLLER] Status of resp value: %+v\n", resp)
		a.logs.ErrorLogger.Printf("[CONTROLLER] Error reason: %+v\n", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Error checkout",
				"status":  http.StatusInternalServerError,
				"data":    resp,
			})
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"message": "Success checkout",
			"status":  http.StatusOK,
			"data":    resp,
		})
}
