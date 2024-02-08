package macaroon

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
)

func TestEncodeMacaroon(t *testing.T) {

	ctx := context.Background()

	abs_path, err := filepath.Abs(".")

	if err != nil {
		t.Fatalf("Failed to derive absolute path, %v", err)
	}

	path_key := filepath.Join(abs_path, "fixtures/signing.key")
	key_uri := fmt.Sprintf("file://%s", path_key)

	loc := "sfomuseum.org"
	duration := "PT1M"

	m_uri, err := NewMacaroonURI(loc, key_uri, duration)

	if err != nil {
		t.Fatalf("Failed to create URI, %v", err)
	}

	m, err := NewMacaroon(ctx, m_uri)

	if err != nil {
		t.Fatalf("Failed to create new macaroon, %v", err)
	}

	enc_m, err := EncodeMacaroonAsBase64(m, false)

	if err != nil {
		t.Fatalf("Failed to encode macaroon, %v", err)
	}

	m2, err := DecodeMacaroonFromBase64(enc_m, false)

	if err != nil {
		t.Fatalf("Failed to decode macaroon, %v", err)
	}

	if IsExpired(m2) {
		t.Fatalf("Macaroon has already expired")
	}

	key, err := NewSigningKey(ctx, key_uri)

	if err != nil {
		t.Fatalf("Failed to retrieve signing key for %s, %v", key_uri, err)
	}

	_, err = m2.Verify(key, [][]byte{}, nil)

	if err != nil {
		t.Fatalf("Macaroon failed to validate")
	}
}
