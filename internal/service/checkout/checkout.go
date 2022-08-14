package checkout

import (
	"context"
	"fmt"

	"github.com/Tiger-Coders/tigerlily-payment/api/rpc"
	"github.com/Tiger-Coders/tigerlily-payment/internal/pkg/logger"
	"github.com/Tiger-Coders/tigerlily-payment/internal/pkg/stripe"
	"github.com/Tiger-Coders/tigerlily-payment/internal/repository/checkout"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Service struct {
	db   *gorm.DB
	base checkout.CheckoutRepo
	logs logger.Logger
	rpc.UnimplementedCheckoutServiceServer
}

var _ rpc.CheckoutServiceServer = (*Service)(nil)

func NewCheckoutService(DB *gorm.DB) *Service {
	return &Service{
		db:   DB,
		base: *checkout.NewCheckoutRepo(DB),
		logs: *logger.NewLogger(),
	}
}

func (srv Service) CustomCheckout(ctx context.Context, req *rpc.CheckoutReq) (resp *rpc.CheckoutResp, err error) {
	fmt.Println("Send request to Notification Service to fire Email, SMS, notification to client and merchant")
	fmt.Printf("THIS IS THE CHECKOUT ITEMS => %+v", req.CheckoutItems)

	checkoutSuccess, err := srv.base.CreateNewOrder(req.CheckoutItems)
	if err != nil {
		srv.logs.ErrorLogger.Printf("[SERVICE] Error processing database transaction: %+v\n", err)
		srv.logs.ErrorLogger.Printf(" [SERVICE] RESULT FROM DS : %+v", checkoutSuccess)
	}

	srv.logs.InfoLogger.Printf(" [SERVICE] CREATE NEW ORDER STATUS : %+v\n", checkoutSuccess)

	// USE ENUM AS ERROR CODES
	resp = &rpc.CheckoutResp{
		Success:   checkoutSuccess,
		StatusUrl: "This is a strategy test.",
		Message:   "Returned from CustomCheckout API; To send a request to notification service for oder confirmation after creating a new order",
	}

	return
}

func (srv Service) StripeCheckoutSession(ctx context.Context, req *rpc.CheckoutReq) (resp *rpc.CheckoutResp, err error) {
	srv.logs.InfoLogger.Printf(" [SERVICE] Checkout service ran %+v", req)

	statusURL := ""

	// TO HANDLE PAYMENT STATUS RESPONSE (RETURN THE STATUS URL BACK TO THE CLIENT)
	statusURL, err = stripe.CreateCheckoutSession()
	if err != nil {
		resp = &rpc.CheckoutResp{
			Success:   false,
			StatusUrl: statusURL,
		}
		return
	}

	checkoutSuccess, err := srv.base.CreateNewOrder(req.CheckoutItems)
	if err != nil {
		srv.logs.ErrorLogger.Printf("[SERVICE] Error processing database transaction: %+v\n", err)
		srv.logs.ErrorLogger.Printf(" [SERVICE] RESULT FROM DS : %+v", checkoutSuccess)
	}

	srv.logs.InfoLogger.Printf(" [SERVICE] CREATE NEW ORDER STATUS : %+v\n", checkoutSuccess)

	// USE ENUM AS ERROR CODES
	resp = &rpc.CheckoutResp{
		Success:   checkoutSuccess,
		StatusUrl: statusURL,
		Message:   "",
	}

	return
}
