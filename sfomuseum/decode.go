package sfomuseum

import (
	"encoding/base64"
	"fmt"

	"github.com/superfly/macaroon"
)

func DecodeMacaroonFromBase64(b64 string) (*macaroon.Macaroon, error) {

	b, err := base64.StdEncoding.DecodeString(b64)

	if err != nil {
		return nil, fmt.Errorf("Failed to decode b2, %v", err)
	}

	return macaroon.Decode(b)
}
