# abacatepay-go-sdk

Official Go SDK for the AbacatePay API.

---

## Installation

```bash
go get github.com/AbacatePay/abacatepay-go-sdk
```

## Quick start

```go
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

	// Create a new billing
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
}
```

# Configuration
| Field	 | Description | Required |
|---------|------------|---------------------|
| APIKey  |	Your AbacatePay API key | yes |
| BaseURL |	Custom API base URL (optional) | no |
| Timeout |	Default HTTP timeout| no |

The SDK also supports configuring the API base URL via environment variable:

```bash
export ABACATEPAY_API_URL=https://api.abacatepay.com
```

### Context & timeouts

All requests require a context.Context.
This allows request cancellation, deadlines and better control in production environments.

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
```

## API Reference

Full API documentation is available at:

https://docs.abacatepay.com/

