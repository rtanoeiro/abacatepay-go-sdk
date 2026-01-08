package main

import (
	"context"
	"log"
	"time"

	"github.com/AbacatePay/abacatepay-go-sdk/abacatepay"
	"github.com/AbacatePay/abacatepay-go-sdk/v1/billing"
)

func main() {
	client, err := abacatepay.New(abacatepay.ClientConfig{
		APIKey:  "abc_dev",
		Timeout: 10 * time.Second,
	})
	
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	body := &billing.CreateBillingBody{
		Frequency:     billing.OneTime,
		Methods:       []billing.Method{billing.PIX},
		CompletionUrl: "https://example.com/completion",
		ReturnUrl:     "https://example.com/return",
		Products: []*billing.BillingProduct{
			{
				ExternalId:  "pix-1234",
				Name:        "Example Product",
				Description: "Example product description",
				Quantity:    1,
				Price:       100,
			},
		},
		Customer: &billing.BillingCustomer{
			Email: "test@example.com",
		},
	}

	createResp, err := client.Billing.Create(ctx, body)
	
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("billing created: %s\n", createResp.Data.PublicID)

	listResp, err := client.Billing.ListAll(ctx)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("total billings: %d\n", len(listResp.Data))
}
