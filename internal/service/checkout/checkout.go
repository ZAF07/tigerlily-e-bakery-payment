package checkout

import (
	"context"

	rpc "github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc/proto"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/manager"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/repository/checkout"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Service struct {
	db *gorm.DB
	base checkout.CheckoutRepo
	logs logger.Logger
	rpc.UnimplementedCheckoutServiceServer
}

var _ rpc.CheckoutServiceServer = (*Service)(nil)

func NewCheckoutService(DB *gorm.DB) *Service {
	return&Service{
		db: DB,
		base: *checkout.NewCheckoutRepo(DB),
		logs: *logger.NewLogger(),
	}
}

func (srv Service) Checkout(ctx context.Context, req *rpc.CheckoutReq) (resp *rpc.CheckoutResp, err error) {
	srv.logs.InfoLogger.Printf(" [SERVICE] Checkout service ran %+v", req)

	// TO HANDLE PAYMENT STATUS RESPONSE (RETURN THE STATUS URL BACK TO THE CLIENT)
	statusURL := manager.CreateCheckoutSession() 

	checkoutSuccess, err := srv.base.CreateNewOrder(req.CheckoutItems)
	if err != nil {
		srv.logs.ErrorLogger.Printf("[SERVICE] Error processing database transaction: %+v\n", err)
		srv.logs.ErrorLogger.Printf(" [SERVICE] RESUKT FROM DS : %+v", checkoutSuccess)
	}

	srv.logs.InfoLogger.Printf(" [SERVICE] RESULT : %+v\n",checkoutSuccess)

	// USE ENUM AS ERROR CODES
	resp = &rpc.CheckoutResp{
		Success: checkoutSuccess,
		StatusUrl: statusURL,
	}

	return 
} 
