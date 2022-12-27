package checkout

import (
	"context"

	"github.com/ZAF07/tigerlily-e-bakery-payment/api/rpc"
)

type CheckoutDBInterface interface {
	CreateNewOrder(ctx context.Context, checkoutItems []*rpc.Checkout) error
}
