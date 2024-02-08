// Generate a random Base64-encoded Macaroon signing key.
package main

import (
	"encoding/base64"
	"fmt"

	"github.com/superfly/macaroon"
)

func main() {

	k := macaroon.NewSigningKey()
	k_b64 := base64.StdEncoding.EncodeToString(k)

	fmt.Println(k_b64)
}
