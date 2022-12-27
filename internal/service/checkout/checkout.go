package checkout

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/pkg/stripe"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/repository/checkout"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Service struct {
	// db   *gorm.DB
	db   *sql.DB
	base checkout.CheckoutRepo
	logs logger.Logger
	rpc.UnimplementedCheckoutServiceServer
}

var _ rpc.CheckoutServiceServer = (*Service)(nil)

func NewCheckoutService(DB *sql.DB) *Service {
	return &Service{
		db:   DB,
		base: *checkout.NewCheckoutRepo(DB),
		logs: *logger.NewLogger(),
	}
}

func (srv Service) CustomCheckout(ctx context.Context, req *rpc.CheckoutReq) (resp *rpc.CheckoutResp, err error) {
	fmt.Println("Send request to Notification Service to fire Email, SMS, notification to client and merchant")
	fmt.Printf("THIS IS THE CHECKOUT ITEMS => %+v\n", req.CheckoutItems)

	err = srv.base.CreateNewOrder(ctx, req.CheckoutItems)
	if err != nil {
		srv.logs.ErrorLogger.Printf("[SERVICE] Error processing database transaction: %+v\n", err)
	}

	srv.logs.InfoLogger.Println(" [SERVICE] CREATE NEW ORDER SUCCESS ✅")

	// USE ENUM AS ERROR CODES
	resp = &rpc.CheckoutResp{
		Success:   true,
		StatusUrl: "This is a strategy test.",
		Message:   "Returned from CustomCheckout API; To send a request to notification service for oder confirmation after creating a new order",
	}

	return
}

func (srv Service) StripeCheckoutSession(ctx context.Context, req *rpc.CheckoutReq) (resp *rpc.CheckoutResp, err error) {
	srv.logs.InfoLogger.Printf(" [SERVICE] Checkout service ran %+v", req)

	statusURL := ""

	// TO HANDLE PAYMENT STATUS RESPONSE (RETURN THE STATUS URL BACK TO THE CLIENT)
	//  TODO : Send amount to charge and item name
	statusURL, err = stripe.CreateCheckoutSession()
	if err != nil {
		resp = &rpc.CheckoutResp{
			Success:   false,
			StatusUrl: statusURL,
		}
		return
	}

	err = srv.base.CreateNewOrder(ctx, req.CheckoutItems)
	if err != nil {
		srv.logs.ErrorLogger.Printf("[SERVICE] Error processing database transaction: %+v\n", err)
	}

	srv.logs.InfoLogger.Printf(" [SERVICE] CREATE NEW ORDER SUCCESS ✅")

	// USE ENUM AS ERROR CODES
	resp = &rpc.CheckoutResp{
		Success:   true,
		StatusUrl: statusURL,
		Message:   "",
	}

	return
}
