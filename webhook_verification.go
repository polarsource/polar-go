package polargo

import (
	"encoding/base64"
	"fmt"
	"net/http"

	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
)

// WebhookVerifier handles verification of incoming Polar webhook signatures.
// Polar uses the Standard Webhooks specification for signing webhook payloads.
type WebhookVerifier struct {
	wh *standardwebhooks.Webhook
}

// NewWebhookVerifier creates a new WebhookVerifier with the provided webhook secret.
// The secret should be the raw webhook secret string from your Polar dashboard.
func NewWebhookVerifier(secret string) (*WebhookVerifier, error) {
	// Standard webhooks expects base64 encoded secret
	encodedSecret := base64.StdEncoding.EncodeToString([]byte(secret))
	wh, err := standardwebhooks.NewWebhook(encodedSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create webhook verifier: %w", err)
	}
	return &WebhookVerifier{wh: wh}, nil
}

// Verify validates the webhook payload against the signature in the request headers.
// It checks the webhook-id, webhook-timestamp, and webhook-signature headers.
// Returns nil if the signature is valid, otherwise returns an error.
func (v *WebhookVerifier) Verify(payload []byte, headers http.Header) error {
	return v.wh.Verify(payload, headers)
}

// ValidateWebhookSignature is a convenience function that creates a verifier and
// validates the webhook signature in a single call.
// The secret should be the raw webhook secret string from your Polar dashboard.
func ValidateWebhookSignature(payload []byte, headers http.Header, secret string) error {
	verifier, err := NewWebhookVerifier(secret)
	if err != nil {
		return err
	}
	return verifier.Verify(payload, headers)
}
