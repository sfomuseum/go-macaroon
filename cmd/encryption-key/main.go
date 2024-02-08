// Generate a random base64-encoded Macaroon encyption key.
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"

	flyio_macaroon "github.com/superfly/macaroon"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Generate a random base64-encoded Macaroon encyption key.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	k := flyio_macaroon.NewEncryptionKey()
	k_b64 := base64.StdEncoding.EncodeToString(k)

	fmt.Println(k_b64)
}
