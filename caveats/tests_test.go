package caveats

// https://pkg.go.dev/github.com/superfly/macaroon
// https://fly.io/blog/macaroons-escalated-quickly/

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	sfom_macaroon "github.com/sfomuseum/go-macaroon"
)

func TestSFOMuseumCaveats(t *testing.T) {

	ctx := context.Background()

	abs_path, err := filepath.Abs(".")

	if err != nil {
		t.Fatalf("Failed to derive absolute path, %v", err)
	}

	path_key := filepath.Join(abs_path, "../fixtures/signing.key")
	key_uri := fmt.Sprintf("file://%s", path_key)

	loc := "sfomuseum.org"
	duration := "PT10S"

	m_uri, err := sfom_macaroon.NewMacaroonURI(loc, key_uri, duration)

	if err != nil {
		t.Fatalf("Failed to create URI, %v", err)
	}

	m, err := sfom_macaroon.NewMacaroon(ctx, m_uri)

	if err != nil {
		t.Fatalf("Failed to create new macaroon, %v", err)
	}

	// Allowed users

	c1 := &isUserCaveat{
		Users: []string{"bob", "alice"},
	}

	// Required role

	c2 := &hasRoleCaveat{
		Role: "staff",
	}

	err = m.Add(c1, c2)

	if err != nil {
		t.Fatalf("Failed to add caveats, %v", err)
	}

	enc, err := sfom_macaroon.EncodeMacaroonAsBase64(m, false)

	if err != nil {
		t.Fatalf("Failed to encode, %v", err)
	}

	fmt.Printf("Macaroon expires %v\n", m.Expiration())
	fmt.Println("Sleeping 9 seconds...")
	time.Sleep(9 * time.Second)

	m2, err := sfom_macaroon.DecodeMacaroonFromBase64(enc, false)

	if err != nil {
		t.Fatalf("Failed to decode, %v", err)
	}

	key, err := sfom_macaroon.NewSigningKey(ctx, key_uri)

	if err != nil {
		t.Fatalf("Failed to retrieve signing key for %s, %v", key_uri, err)
	}

	cs, err := m2.Verify(key, [][]byte{}, nil)

	if err != nil {
		t.Fatalf("Failed to verify, %v", err)
	}

	err = cs.Validate(
		&isUserAccess{User: "alice"},
		&hasRoleAccess{Role: "staff"},
	)

	if err != nil {
		t.Fatalf("Failed to validate alice, %v", err)
	}

	fmt.Println("OK alice")

	err = cs.Validate(
		&isUserAccess{User: "bob"},
		&hasRoleAccess{Role: "staff"},
	)

	if err != nil {
		t.Fatalf("Failed to validate bob, %v", err)
	}

	fmt.Println("OK bob")

	err = cs.Validate(
		&isUserAccess{User: "doug"},
		&hasRoleAccess{Role: "staff"},
	)

	if err == nil {
		t.Fatalf("NO DOUG FOR YOU")
	}

	fmt.Println("DENY doug (this is good)")

	fmt.Println("Sleep again 2 seconds...")
	time.Sleep(2 * time.Second)

	err = cs.Validate(
		&isUserAccess{User: "bob"},
		&hasRoleAccess{Role: "staff"},
	)

	if err == nil {
		t.Fatalf("Macaroon should have expired")
	}

	fmt.Printf("Failed to validate bob, %v (this is good)\n", err)

	if !sfom_macaroon.IsExpired(m2) {
		t.Fatalf("Macaroon should be expired")
	}
}
