package caveats

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	sfom_macaroon "github.com/sfomuseum/go-macaroon"
)

func TestEnsureAccount(t *testing.T) {

	ctx := context.Background()

	abs_path, err := filepath.Abs(".")

	if err != nil {
		t.Fatalf("Failed to derive absolute path, %v", err)
	}

	path_key := filepath.Join(abs_path, "../fixtures/signing.key")
	key_uri := fmt.Sprintf("file://%s", path_key)

	key, err := sfom_macaroon.NewSigningKey(ctx, key_uri)

	if err != nil {
		t.Fatalf("Failed to retrieve signing key for %s, %v", key_uri, err)
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

	c := &EnsureAccountCaveat{
		AccountId: 1,
	}

	err = m.Add(c)

	if err != nil {
		t.Fatalf("Failed to add caveats, %v", err)
	}

	enc, err := sfom_macaroon.EncodeMacaroonAsBase64(m, false)

	if err != nil {
		t.Fatalf("Failed to encode, %v", err)
	}

	m2, err := sfom_macaroon.DecodeMacaroonFromBase64(enc, false)

	if err != nil {
		t.Fatalf("Failed to decode, %v", err)
	}

	cs, err := m2.Verify(key, [][]byte{}, nil)

	if err != nil {
		t.Fatalf("Failed to verify, %v", err)
	}

	err = cs.Validate(
		&EnsureAccountAccess{AccountId: 1},
	)

	if err != nil {
		t.Fatalf("Expected the access to validate, %v", err)
	}

	err = cs.Validate(
		&EnsureAccountAccess{AccountId: 2},
	)

	if err == nil {
		t.Fatalf("Expected the access NOT to validate")
	}
}

func TestEnsureAccountWithRolesAll(t *testing.T) {

	ctx := context.Background()

	abs_path, err := filepath.Abs(".")

	if err != nil {
		t.Fatalf("Failed to derive absolute path, %v", err)
	}

	path_key := filepath.Join(abs_path, "../fixtures/signing.key")
	key_uri := fmt.Sprintf("file://%s", path_key)

	key, err := sfom_macaroon.NewSigningKey(ctx, key_uri)

	if err != nil {
		t.Fatalf("Failed to retrieve signing key for %s, %v", key_uri, err)
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

	c := &EnsureAccountCaveat{
		AccountId: 1,
		RolesAll:  []string{"staff", "admin"},
	}

	err = m.Add(c)

	if err != nil {
		t.Fatalf("Failed to add caveats, %v", err)
	}

	enc, err := sfom_macaroon.EncodeMacaroonAsBase64(m, false)

	if err != nil {
		t.Fatalf("Failed to encode, %v", err)
	}

	m2, err := sfom_macaroon.DecodeMacaroonFromBase64(enc, false)

	if err != nil {
		t.Fatalf("Failed to decode, %v", err)
	}

	cs, err := m2.Verify(key, [][]byte{}, nil)

	if err != nil {
		t.Fatalf("Failed to verify, %v", err)
	}

	err = cs.Validate(
		&EnsureAccountAccess{
			AccountId: 1,
			Roles:     []string{"staff", "admin"},
		},
	)

	if err != nil {
		t.Fatalf("Expected the access to validate, %v", err)
	}

	err = cs.Validate(
		&EnsureAccountAccess{
			AccountId: 2,
		},
	)

	if err == nil {
		t.Fatalf("Expected the access NOT to validate (invalid account)")
	}

	err = cs.Validate(
		&EnsureAccountAccess{
			AccountId: 1,
			Roles:     []string{"staff"},
		},
	)

	if err == nil {
		t.Fatalf("Expected the access NOT to validate (insufficient roles)")
	}
}

func TestEnsureAccountWithRolesAnyl(t *testing.T) {

	ctx := context.Background()

	abs_path, err := filepath.Abs(".")

	if err != nil {
		t.Fatalf("Failed to derive absolute path, %v", err)
	}

	path_key := filepath.Join(abs_path, "../fixtures/signing.key")
	key_uri := fmt.Sprintf("file://%s", path_key)

	key, err := sfom_macaroon.NewSigningKey(ctx, key_uri)

	if err != nil {
		t.Fatalf("Failed to retrieve signing key for %s, %v", key_uri, err)
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

	c := &EnsureAccountCaveat{
		AccountId: 1,
		RolesAny:  []string{"staff", "admin"},
	}

	err = m.Add(c)

	if err != nil {
		t.Fatalf("Failed to add caveats, %v", err)
	}

	enc, err := sfom_macaroon.EncodeMacaroonAsBase64(m, false)

	if err != nil {
		t.Fatalf("Failed to encode, %v", err)
	}

	m2, err := sfom_macaroon.DecodeMacaroonFromBase64(enc, false)

	if err != nil {
		t.Fatalf("Failed to decode, %v", err)
	}

	cs, err := m2.Verify(key, [][]byte{}, nil)

	if err != nil {
		t.Fatalf("Failed to verify, %v", err)
	}

	err = cs.Validate(
		&EnsureAccountAccess{
			AccountId: 1,
			Roles:     []string{"staff", "admin"},
		},
	)

	if err != nil {
		t.Fatalf("Expected the access to validate, %v", err)
	}

	err = cs.Validate(
		&EnsureAccountAccess{
			AccountId: 2,
		},
	)

	if err == nil {
		t.Fatalf("Expected the access NOT to validate (invalid account)")
	}

	err = cs.Validate(
		&EnsureAccountAccess{
			AccountId: 1,
			Roles:     []string{"staff"},
		},
	)

	if err != nil {
		t.Fatalf("Expected the access to validate (has at least one role), %v", err)
	}
}
