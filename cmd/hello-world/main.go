package main

// https://pkg.go.dev/github.com/superfly/macaroon
// https://fly.io/blog/macaroons-escalated-quickly/

import (
	"log"
	"time"

	"github.com/sfomuseum/go-macaroon/sfomuseum"
	"github.com/superfly/macaroon"
)

func main() {

	// Random every time
	key := macaroon.NewSigningKey()

	m, err := sfomuseum.TestMacaroon(key)

	if err != nil {
		log.Fatalf("Failed to create test macaroon, %v", err)
	}

	enc, err := sfomuseum.EncodeMacaroonAsBase64(m)

	if err != nil {
		log.Fatalf("Failed to encode, %v", err)
	}

	log.Println("Expires", m.Expiration())
	log.Println("Sleeping...")
	time.Sleep(9 * time.Second)

	m2, err := sfomuseum.DecodeMacaroonFromBase64(enc)

	if err != nil {
		log.Fatalf("Failed to decode, %v", err)
	}

	cs, err := m2.Verify(key, [][]byte{}, nil)

	if err != nil {
		log.Fatalf("Failed to verify, %v", err)
	}

	err = cs.Validate(
		&sfomuseum.IsUserAccess{User: "alice"},
		&sfomuseum.HasRoleAccess{Role: "staff"},
	)

	if err != nil {
		log.Fatalf("Failed to validate alice, %v", err)
	}

	log.Println("OK alice")

	err = cs.Validate(
		&sfomuseum.IsUserAccess{User: "bob"},
		&sfomuseum.HasRoleAccess{Role: "staff"},
	)

	if err != nil {
		log.Fatalf("Failed to validate bob, %v", err)
	}

	log.Println("OK bob")

	err = cs.Validate(
		&sfomuseum.IsUserAccess{User: "doug"},
		&sfomuseum.HasRoleAccess{Role: "staff"},
	)

	if err == nil {
		log.Fatalf("NO DOUG FOR YOU")
	}

	log.Println("DENY doug")

	log.Println("Sleep again...")
	time.Sleep(2 * time.Second)

	err = cs.Validate(
		&sfomuseum.IsUserAccess{User: "bob"},
		&sfomuseum.HasRoleAccess{Role: "staff"},
	)

	if err == nil {
		log.Fatalf("SHOULD HAVE EXPIRED")
	}

	log.Printf("Failed to validate bob, %v\n", err)
}
