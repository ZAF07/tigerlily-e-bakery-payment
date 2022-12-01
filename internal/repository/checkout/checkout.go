package checkout

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/pkg/logger"

	// _ "github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type CheckoutRepo struct {
	db   *sql.DB
	logs logger.Logger
}

func NewCheckoutRepo(DB *sql.DB) *CheckoutRepo {
	return &CheckoutRepo{
		db:   DB,
		logs: *logger.NewLogger(),
	}
}

func (repo CheckoutRepo) CreateNewOrder(ctx context.Context, checkoutItems []*rpc.Checkout) error {

	// Get a connection from the pool
	// db, err := injection.GetPaymentDBInstance().Conn(ctx)
	db, err := repo.db.Conn(ctx)
	if err != nil {
		log.Fatalf("Error getting connection from DB pool: %+v", err)
		return err
	}

	//  Begin TX with default options, (💡 Might want to check out the options)
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		log.Printf("database transaction error : %+v\n", err)
		return err
	}
	for _, v := range checkoutItems {
		now := time.Now()
		_, eErr := tx.Exec("INSERT INTO orders (order_id, sku_id, customer_id, discount_code, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", v.OrderId, v.SkuId, v.CustomerId, v.DiscountCode, now, now)
		if eErr != nil {
			log.Printf("EXEC ERROR --> %+v", eErr)
			tx.Rollback()
			return eErr
		}
	}

	if cErr := tx.Commit(); cErr != nil {
		tx.Rollback()
		log.Printf("error committing transaction : %+v\n", cErr)
		return cErr
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

	return err
}
