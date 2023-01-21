package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

func Payment(products string, totalAmount float64) {
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
		SuccessURL: stripe.String("http://localhost:3001/utils/success.html"),
		CancelURL:  stripe.String("http://localhost:3001/utils/cancel.html"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}
	fmt.Println(fiber.Map{"message": "success",
		"products": "orders",
		"total":    total,
		"url":      s.URL,
	})
}
