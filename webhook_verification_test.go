package polargo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testSecret = "whsec_test_secret_key_for_testing"

// generateTestSignature creates a valid webhook signature for testing
func generateTestSignature(payload []byte, secret string, timestamp int64, msgID string) string {
	// Standard webhooks signature format: v1,<base64-hmac-sha256>
	toSign := fmt.Sprintf("%s.%d.%s", msgID, timestamp, string(payload))

	// Decode the base64 secret (standard webhooks expects base64 encoded secret)
	encodedSecret := base64.StdEncoding.EncodeToString([]byte(secret))
	decodedSecret, _ := base64.StdEncoding.DecodeString(encodedSecret)

	h := hmac.New(sha256.New, decodedSecret)
	h.Write([]byte(toSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return "v1," + signature
}

func TestNewWebhookVerifier(t *testing.T) {
	t.Run("creates verifier with valid secret", func(t *testing.T) {
		verifier, err := NewWebhookVerifier(testSecret)
		require.NoError(t, err)
		assert.NotNil(t, verifier)
	})

	t.Run("creates verifier with empty secret", func(t *testing.T) {
		// Empty secret should still create a verifier (validation happens at verify time)
		verifier, err := NewWebhookVerifier("")
		require.NoError(t, err)
		assert.NotNil(t, verifier)
	})
}

func TestWebhookVerifier_Verify(t *testing.T) {
	payload := []byte(`{"type":"customer.updated","data":{"id":"123"}}`)
	timestamp := time.Now().Unix()
	msgID := "msg_test123"

	t.Run("valid signature passes verification", func(t *testing.T) {
		verifier, err := NewWebhookVerifier(testSecret)
		require.NoError(t, err)

		signature := generateTestSignature(payload, testSecret, timestamp, msgID)

		headers := http.Header{}
		headers.Set("webhook-id", msgID)
		headers.Set("webhook-timestamp", strconv.FormatInt(timestamp, 10))
		headers.Set("webhook-signature", signature)

		err = verifier.Verify(payload, headers)
		assert.NoError(t, err)
	})

	t.Run("invalid signature fails verification", func(t *testing.T) {
		verifier, err := NewWebhookVerifier(testSecret)
		require.NoError(t, err)

		headers := http.Header{}
		headers.Set("webhook-id", msgID)
		headers.Set("webhook-timestamp", strconv.FormatInt(timestamp, 10))
		headers.Set("webhook-signature", "v1,invalidsignature")

		err = verifier.Verify(payload, headers)
		assert.Error(t, err)
	})

	t.Run("missing headers fails verification", func(t *testing.T) {
		verifier, err := NewWebhookVerifier(testSecret)
		require.NoError(t, err)

		headers := http.Header{}
		// Missing all required headers

		err = verifier.Verify(payload, headers)
		assert.Error(t, err)
	})

	t.Run("wrong secret fails verification", func(t *testing.T) {
		verifier, err := NewWebhookVerifier("wrong_secret")
		require.NoError(t, err)

		signature := generateTestSignature(payload, testSecret, timestamp, msgID)

		headers := http.Header{}
		headers.Set("webhook-id", msgID)
		headers.Set("webhook-timestamp", strconv.FormatInt(timestamp, 10))
		headers.Set("webhook-signature", signature)

		err = verifier.Verify(payload, headers)
		assert.Error(t, err)
	})
}

func TestValidateWebhookSignature(t *testing.T) {
	payload := []byte(`{"type":"subscription.active","data":{"id":"sub_123"}}`)
	timestamp := time.Now().Unix()
	msgID := "msg_sub_test"

	t.Run("convenience function works with valid signature", func(t *testing.T) {
		signature := generateTestSignature(payload, testSecret, timestamp, msgID)

		headers := http.Header{}
		headers.Set("webhook-id", msgID)
		headers.Set("webhook-timestamp", strconv.FormatInt(timestamp, 10))
		headers.Set("webhook-signature", signature)

		err := ValidateWebhookSignature(payload, headers, testSecret)
		assert.NoError(t, err)
	})

	t.Run("convenience function fails with invalid signature", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("webhook-id", msgID)
		headers.Set("webhook-timestamp", strconv.FormatInt(timestamp, 10))
		headers.Set("webhook-signature", "v1,bad_signature")

		err := ValidateWebhookSignature(payload, headers, testSecret)
		assert.Error(t, err)
	})
}
