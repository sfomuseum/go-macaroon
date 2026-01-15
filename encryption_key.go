package macaroon

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/aaronland/gocloud/runtimevar"
	flyio_macaroon "github.com/superfly/macaroon"
)

func NewEncryptionKey(ctx context.Context, key_uri string) (flyio_macaroon.EncryptionKey, error) {

	b64_key, err := runtimevar.StringVar(ctx, key_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive key, %w", err)
	}

	key, err := base64.StdEncoding.DecodeString(b64_key)

	if err != nil {
		return nil, fmt.Errorf("Failed to decod key, %v", err)
	}

	return flyio_macaroon.EncryptionKey(key), nil
}
