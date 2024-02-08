package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"

	sfom_macaroon "github.com/sfomuseum/go-macaroon"
)

func main() {

	var key_uri string
	var loc string
	var duration string
	var urlescape bool

	flag.StringVar(&key_uri, "signing-key-uri", "", "...")
	flag.StringVar(&loc, "location", "sfomuseum.org", "...")
	flag.StringVar(&duration, "duration", "PT10M", "...")
	flag.BoolVar(&urlescape, "urlescape", false, "...")

	flag.Parse()

	ctx := context.Background()

	m_params := &url.Values{}
	m_params.Set("key", key_uri)
	m_params.Set("duration", duration)

	m_uri := url.URL{}
	m_uri.Scheme = "macaroon"
	m_uri.Host = loc
	m_uri.RawQuery = m_params.Encode()

	m, err := sfom_macaroon.NewMacaroon(ctx, m_uri.String())

	if err != nil {
		log.Fatalf("Failed to create macaroon, %v", err)
	}

	enc, err := sfom_macaroon.EncodeMacaroonAsBase64(m, urlescape)

	if err != nil {
		log.Fatalf("Failed to encode, %v", err)
	}

	fmt.Println(enc)
}
