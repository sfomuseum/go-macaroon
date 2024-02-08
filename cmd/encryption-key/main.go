// Generate a random Base64-encoded Macaroon encyption key.
package main

import (
	"encoding/base64"
	"fmt"

	flyio_macaroon "github.com/superfly/macaroon"
)

func main() {

	k := flyio_macaroon.NewEncryptionKey()
	k_b64 := base64.StdEncoding.EncodeToString(k)

	fmt.Println(k_b64)
}
