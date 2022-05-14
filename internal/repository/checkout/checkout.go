package checkout

import (
	rpc "github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc/proto"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/models"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/pkg/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type CheckoutRepo struct {
	db *gorm.DB
	logs logger.Logger
}

func NewCheckoutRepo(DB *gorm.DB) *CheckoutRepo {
	return&CheckoutRepo{
		db: DB,
		logs: *logger.NewLogger(),
	}
}

func (repo CheckoutRepo) CreateNewOrder(checkoutItems []*rpc.Checkout) (success bool, err error) {

	// RUN A TRANSACTION FOR CREATION, IF FAIL, WILL FALLBACK
	repo.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range checkoutItems {
			orderItem := &models.Order{ // Should add price into table
				DiscountCode: item.DiscountCode,
				OrderID: item.OrderId,
				CustomerID: item.CustomerId,
				SkuID: item.SkuId,
			}
	
			if err := tx.Debug().Omit("DeletedAt").Create(&orderItem).Error; err != nil {
				repo.logs.WarnLogger.Printf("[REPO] Error batch creating order items : %+v", err)
				success = false
				return err
			}
		}
		success = true
		return nil
	})
	
	return
}