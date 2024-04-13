package checkout

import (
	"context"

	"github.com/Tiger-Coders/tigerlily-payment/api/rpc"
)

type CheckoutDBInterface interface {
	CreateNewOrder(ctx context.Context, checkoutItems []*rpc.Checkout) error
	Ping() error
}
