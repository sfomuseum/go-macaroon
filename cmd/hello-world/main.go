package main

// https://pkg.go.dev/github.com/superfly/macaroon
// https://fly.io/blog/macaroons-escalated-quickly/

import (
	"encoding/base64"
	"log"
	"time"

	"github.com/sfomuseum/go-macaroon/sfomuseum"
	"github.com/superfly/macaroon"
)

func main() {

	kid := make([]byte, 0)
	loc := "sfomuseum.org"

	// Random every time
	key := macaroon.NewSigningKey()

	m, err := macaroon.New(kid, loc, key)

	if err != nil {
		log.Fatalf("Failed to create macaroon, %v", err)
	}

	// Max lifetime

	max_d := 10 * time.Second

	c := &macaroon.ValidityWindow{
		NotBefore: time.Now().Unix(),
		NotAfter:  time.Now().Add(max_d).Unix(),
	}

	// Allowed users

	c2 := &sfomuseum.IsUserCaveat{
		Users: []string{"bob", "alice"},
	}

	// Required role

	c3 := &sfomuseum.HasRoleCaveat{
		Role: "staff",
	}

	err = m.Add(c, c2, c3)

	if err != nil {
		log.Fatalf("Failed to add caveat, %v", err)
	}

	b, err := m.Encode()

	if err != nil {
		log.Fatalf("Failed to encode, %v", err)
	}

	s := base64.StdEncoding.EncodeToString(b)

	b2, err := base64.StdEncoding.DecodeString(s)

	if err != nil {
		log.Fatalf("Failed to decode b2, %v", err)
	}

	log.Println("Expires", m.Expiration())
	log.Println("Sleeping...")
	time.Sleep(9 * time.Second)

	m2, err := macaroon.Decode(b2)

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
