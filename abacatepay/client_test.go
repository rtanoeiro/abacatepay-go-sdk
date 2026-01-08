package abacatepay_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/AbacatePay/abacatepay-go-sdk/abacatepay"
)

func TestNew(t *testing.T) {
	t.Run("create client with valid config", func(t *testing.T) {
		client, err := abacatepay.New(abacatepay.ClientConfig{
			APIKey:  "test-key",
			Timeout: 10 * time.Second,
		})

		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.NotNil(t, client.Billing)
	})

	t.Run("error when api key is empty", func(t *testing.T) {
		client, err := abacatepay.New(abacatepay.ClientConfig{
			APIKey: "",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, abacatepay.ErrInvalidAPIKey)
		assert.Nil(t, client)
	})
}
