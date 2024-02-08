package macaroon

import (
	"encoding/base64"
	"fmt"
	"net/url"

	flyio_macaroon "github.com/superfly/macaroon"
)

func DecodeMacaroonFromBase64(b64 string, urlunescape bool) (*flyio_macaroon.Macaroon, error) {

	if urlunescape {
		v, err := url.QueryUnescape(b64)

		if err != nil {
			return nil, fmt.Errorf("Failed to unescape value, %w", err)
		}

		b64 = v
	}

	b, err := base64.StdEncoding.DecodeString(b64)

	if err != nil {
		return nil, fmt.Errorf("Failed to decode b2, %v", err)
	}

	return flyio_macaroon.Decode(b)
}
