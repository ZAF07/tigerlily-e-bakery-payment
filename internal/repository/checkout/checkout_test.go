package checkout_test

import (
	"context"
	"testing"

	"github.com/Tiger-Coders/tigerlily-payment/api/rpc"
	"github.com/Tiger-Coders/tigerlily-payment/internal/db"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	orderItem := &rpc.Checkout{
		OrderId:      "orderId",
		SkuId:        "skuid",
		CustomerId:   "customerId",
		DiscountCode: "discountcode",
	}

	orderItems := []*rpc.Checkout{orderItem}
	db := db.InitPostgresDB()
	db.CreateNewOrder(ctx, orderItems)
}
