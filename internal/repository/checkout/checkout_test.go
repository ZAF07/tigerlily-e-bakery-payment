package checkout_test

import (
	"testing"

	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/db"
	"github.com/ZAF07/tigerlily-e-bakery-payment/internal/models"
)


func TestCreate(t *testing.T) {
	orderItems := &models.Order{
		OrderID: "orderId",
		SkuID: "skuid",
		CustomerID: "customerId",
		DiscountCode: "discountcode",
	}
	db := db.NewDB()
	db.Create(orderItems)

}