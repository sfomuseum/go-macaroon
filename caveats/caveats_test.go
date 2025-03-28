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
	flyio_macaroon "github.com/superfly/macaroon"
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

func TestSFOMuseumThirdPartyCaveats(t *testing.T) {

	ctx := context.Background()

	abs_path, err := filepath.Abs(".")

	if err != nil {
		t.Fatalf("Failed to derive absolute path, %v", err)
	}

	path_key := filepath.Join(abs_path, "../fixtures/signing.key")
	key_uri := fmt.Sprintf("file://%s", path_key)

	path_encryption_key := filepath.Join(abs_path, "../fixtures/3p-encryption.key")
	encryption_key_uri := fmt.Sprintf("file://%s", path_encryption_key)

	key_ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	fmt.Printf("Load sigining key, %s\n", key_uri)

	key, err := sfom_macaroon.NewSigningKey(key_ctx, key_uri)

	if err != nil {
		t.Fatalf("Failed to retrieve signing key for %s, %v", key_uri, err)
	}

	fmt.Printf("Load encryption key, %s\n", encryption_key_uri)

	shared_key, err := sfom_macaroon.NewEncryptionKey(key_ctx, encryption_key_uri)

	if err != nil {
		t.Fatalf("Failed to load encyption key from %s, %v", encryption_key_uri, err)
	}

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

	tp_c := &dayOfWeekCaveat{
		Days: []string{"Friday"},
	}

	tp_loc := "example.com"

	err = m.Add3P(shared_key, tp_loc, tp_c)

	if err != nil {
		t.Fatalf("Failed to add 3P, %v", err)
	}

	// END OF add third-party caveat here

	enc, err := sfom_macaroon.EncodeMacaroonAsBase64(m, false)

	if err != nil {
		t.Fatalf("Failed to encode macaroon, %v", err)
	}

	// Pretend we are sending enc somewhere...

	fmt.Println("Sleeping 2 seconds to simulate exchanging macaroon...")
	time.Sleep(2 * time.Second)

	// Pretend enc is being received somewhere...

	m2, err := sfom_macaroon.DecodeMacaroonFromBase64(enc, false)

	if err != nil {
		t.Fatalf("Failed to decode macaroon, %v", err)
	}

	// START OF third-party discharges

	discharges := make([][]byte, 0)
	discharge_keys := make(map[string][]flyio_macaroon.EncryptionKey)

	tp, err := m2.ThirdPartyTicket(tp_loc)

	if err != nil {
		t.Fatalf("Failed to get third party tickets, %v", err)
	}

	// Pretend we are sending tp to loc here...

	fmt.Println("Sleeping 2 seconds to simulate sending discharge ticket...")
	time.Sleep(2 * time.Second)

	tp_caveats, tp_discharge, err := flyio_macaroon.DischargeTicket(shared_key, tp_loc, tp)

	if err != nil {
		t.Fatalf("Failed to parse discharge ticket, %v", err)
	}

	tp_cs := flyio_macaroon.NewCaveatSet(tp_caveats...)

	err = tp_cs.Validate(
		&dayOfWeekAccess{Day: "Friday"},
	)

	if err != nil {
		// Pretend tp_loc is sending back an error here
		t.Fatalf("TP dispatch failed to validate CS, %v", err)
	}

	encd, err := tp_discharge.Encode()

	if err != nil {
		t.Fatalf("Failed to create discharge token, %v", err)
	}

	// Pretend tp_loc is sending back 'encd' here...

	fmt.Println("Sleeping 2 seconds to simulate receiving discharge token...")
	time.Sleep(2 * time.Second)

	discharges = append(discharges, encd)
	discharge_keys[tp_loc] = []flyio_macaroon.EncryptionKey{
		shared_key,
	}

	// END OF third-party discharges

	// Verify with all the other caveats

	cs, err := m2.Verify(key, discharges, discharge_keys)

	if err != nil {
		t.Fatalf("Failed to verify macaroon, %v", err)
	}

	// Test caveats (defined in sfomuseum.TestMacaroon)

	err = cs.Validate(
		&isUserAccess{User: "alice"},
		&hasRoleAccess{Role: "staff"},
	)

	if err != nil {
		t.Fatalf("Failed to validate alice, %v", err)
	}

	fmt.Println("All caveats validate")

	fmt.Println("Sleeping 4 seconds to trigger expiration")
	time.Sleep(4 * time.Second)

	err = cs.Validate(
		&isUserAccess{User: "alice"},
		&hasRoleAccess{Role: "staff"},
	)

	if err == nil {
		t.Fatalf("Expected macaroon to be expired")
	}

	if !sfom_macaroon.IsExpired(m2) {
		t.Fatalf("Expected macaroon to be expired")
	}

	fmt.Printf("Macaroon expired, %v", err)
}
