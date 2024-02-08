// Generate a new Macaroon token and emit it as a base64-encoded string.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	sfom_macaroon "github.com/sfomuseum/go-macaroon"
)

func main() {

	var key_uri string
	var loc string
	var duration string
	var urlescape bool

	flag.StringVar(&key_uri, "signing-key-uri", "", "A valid sfomuseum/runtimevar URI that contains the signing key for the Macaroon token.")
	flag.StringVar(&loc, "location", "sfomuseum.org", "The 'location' string to associate with the Macaroon token.")
	flag.StringVar(&duration, "duration", "PT10M", "A valid ISO8061 duration time used to set the time-to-live for the Macaroon token.")
	flag.BoolVar(&urlescape, "urlescape", false, "A boolean flag to URL escape the final base64-encoded Macaroon token.")

	flag.Parse()

	ctx := context.Background()

	m_uri, err := sfom_macaroon.NewMacaroonURI(loc, key_uri, duration)

	if err != nil {
		log.Fatalf("Failed to derive new macaroon URI, %v", err)
	}
	
	m, err := sfom_macaroon.NewMacaroon(ctx, m_uri)

	if err != nil {
		log.Fatalf("Failed to create macaroon, %v", err)
	}

	enc, err := sfom_macaroon.EncodeMacaroonAsBase64(m, urlescape)

	if err != nil {
		log.Fatalf("Failed to encode, %v", err)
	}

	fmt.Println(enc)
}
