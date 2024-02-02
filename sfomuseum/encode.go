package sfomuseum

import (
	"encoding/base64"
	"fmt"

	"github.com/superfly/macaroon"
)

func EncodeMacaroonAsBase64(m *macaroon.Macaroon) (string, error) {

	b, err := m.Encode()

	if err != nil {
		return "", fmt.Errorf("Failed to encode, %v", err)
	}

	b64 := base64.StdEncoding.EncodeToString(b)
	return b64, nil
}
