package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	sfom_macaroon "github.com/sfomuseum/go-macaroon"
	sfom_caveats "github.com/sfomuseum/go-macaroon/caveats"
)

func main() {

	var encryption_key_uri string
	var loc string
	var account_id int64
	var b64 string
	var urlunescape bool
	var urlescape bool

	flag.StringVar(&encryption_key_uri, "encryption-key-uri", "", "...")
	flag.StringVar(&loc, "location", "login.sfomuseum", "...")
	flag.Int64Var(&account_id, "account-id", 0, "...")
	flag.StringVar(&b64, "macaroon", "", "...")
	flag.BoolVar(&urlescape, "urlescape", false, "...")
	flag.BoolVar(&urlunescape, "urlunescape", false, "...")

	flag.Parse()

	ctx := context.Background()

	if b64 == "-" {

		b64 = ""
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			b64 += scanner.Text()
		}

		err := scanner.Err()

		if err != nil {
			log.Fatalf("Failed to read from STDIN, %v", err)
		}

	}

	m, err := sfom_macaroon.DecodeMacaroonFromBase64(b64, urlunescape)

	if err != nil {
		log.Fatalf("Failed to base64 decode macaroon, %v", err)
	}

	if sfom_macaroon.IsExpired(m) {
		log.Fatalf("Macaroon has expired")
	}

	shared_k, err := sfom_macaroon.NewEncryptionKey(ctx, encryption_key_uri)

	if err != nil {
		log.Fatalf("Failed to create encryption key, %v", err)
	}

	c := &sfom_caveats.EnsureAccountCaveat{
		AccountId: account_id,
	}

	err = m.Add3P(shared_k, loc, c)

	if err != nil {
		log.Fatalf("Failed to add ensure account caveat, %v", err)
	}

	b64, err = sfom_macaroon.EncodeMacaroonAsBase64(m, urlescape)

	if err != nil {
		log.Fatalf("Failed to encode, %v", err)
	}

	fmt.Println(b64)
}
