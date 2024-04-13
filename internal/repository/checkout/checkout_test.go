package checkout_test

import (
	"testing"

	"github.com/Tiger-Coders/tigerlily-payment/internal/db"
	"github.com/Tiger-Coders/tigerlily-payment/internal/models"
)

// TODO: Complete test case
func TestCreate(t *testing.T) {
	orderItems := &models.Order{
		OrderID:      "orderId",
		SkuID:        "skuid",
		CustomerID:   "customerId",
		DiscountCode: "discountcode",
	}
	db := db.NewDB()
	db.Create(orderItems)
}
