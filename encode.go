package macaroon

import (
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/superfly/macaroon"
)

func EncodeMacaroonAsBase64(m *macaroon.Macaroon, urlescape bool) (string, error) {

	b, err := m.Encode()

	if err != nil {
		return "", fmt.Errorf("Failed to encode, %v", err)
	}

	b64 := base64.StdEncoding.EncodeToString(b)

	if urlescape {
		b64 = url.QueryEscape(b64)
	}

	return b64, nil
}
