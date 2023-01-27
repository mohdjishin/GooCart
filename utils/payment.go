package utils

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

func Payment(products string, totalAmount float64) string {
	total := totalAmount
	stripe.Key = os.Getenv("PAYMENT_SEC_KEY")
	params := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("inr"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(products),
					},
					UnitAmount: stripe.Int64(int64(total * 100)),
				},
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL:        stripe.String("http://localhost:3001/user/order_from_cart"),
		CancelURL:         stripe.String("http://localhost:3001/utils/cancel.html"),
		ClientReferenceID: stripe.String("1"),
	}

	s, err := session.New(params)
	fmt.Println(s.PaymentStatus)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	return s.URL
}
