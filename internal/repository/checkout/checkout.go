package checkout

import (
	"log"
	"time"

	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/injection"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/pkg/logger"
	"github.com/jinzhu/gorm"

	// _ "github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type CheckoutRepo struct {
	db   *gorm.DB
	logs logger.Logger
}

func NewCheckoutRepo(DB *gorm.DB) *CheckoutRepo {
	return &CheckoutRepo{
		db:   DB,
		logs: *logger.NewLogger(),
	}
}

func (repo CheckoutRepo) CreateNewOrder(checkoutItems []*rpc.Checkout) (success bool, err error) {

	//  DI
	db := injection.GetPaymentDBInstance()

	log.Printf("GOT IN REPO : %+v", checkoutItems)

	tx, terr := db.Begin()
	if terr != nil {
		log.Printf("database transaction error : %+v\n", err)
		return false, terr
	}
	for _, v := range checkoutItems {
		now := time.Now()
		_, e := tx.Exec("INSERT INTO orders (order_id, sku_id, customer_id, discount_code, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", v.OrderId, v.SkuId, v.CustomerId, v.DiscountCode, now, now)
		if e != nil {
			log.Printf("EXEC ERROR --> %+v", e)
			tx.Rollback()
			return false, e
		}
	}

	if err = tx.Commit(); err != nil {
		log.Printf("error committing transaction : %+v\n", err)
		return false, err
	}

	// RUN A TRANSACTION FOR CREATION, IF FAIL, WILL FALLBACK
	// repo.db.Transaction(func(tx *gorm.DB) error {
	// 	for _, item := range checkoutItems {
	// 		orderItem := &models.Order{ // Should add price into table
	// 			DiscountCode: item.DiscountCode,
	// 			OrderID:      item.OrderId,
	// 			CustomerID:   item.CustomerId,
	// 			SkuID:        item.SkuId,
	// 		}

	// 		if err := tx.Debug().Omit("DeletedAt").Create(&orderItem).Error; err != nil {
	// 			repo.logs.WarnLogger.Printf("[REPO] Error batch creating order items : %+v", err)
	// 			success = false
	// 			return err
	// 		}
	// 	}
	// 	success = true
	// 	return nil
	// })

	return
}
