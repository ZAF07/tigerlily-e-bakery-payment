package manager

import (
	"fmt"
	"log"

	"github.com/Tiger-Coders/tigerlily-payment/internal/injection"
	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func CreateCheckoutSession() (string, error) {
	config := injection.GetStripeServiceConfig()
	// stripe.Key = "sk_test_51J3JViF3ZrLzHIuoeso9urc2F8T2YSBDcs0Fv3vyHKOq9soeNQxvHP64w8jLZyggHXfJBflHLcFWN3mIRQ2gaOig00iqJFspSK"
	// domain := "http://localhost:3000"
	stripe.Key = config.Authkey
	domain := config.Domain
	currency := config.Currency
	successURL := config.SuccessURL
	cancelURL := config.CancelURL

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				// Price: stripe.String("price_1KzE25F3ZrLzHIuo10cYln0C"),
				Quantity: stripe.Int64(1),
				Currency: stripe.String(currency),
				Amount:   stripe.Int64(1999),
				Name:     stripe.String("Test cheese cake"),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		// SuccessURL: stripe.String(domain + "?success=true"),
		// CancelURL: stripe.String(domain + "?canceled=true"),
		SuccessURL: stripe.String(domain + successURL),
		CancelURL:  stripe.String(domain + cancelURL),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
		return "", err
	}

	fmt.Printf("THIS IS THE RESULTS FROM STRIPE: %+v", s.URL)
	return s.URL, nil
}
