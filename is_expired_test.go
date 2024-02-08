package macaroon

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
)

func TestIsExpires(t *testing.T) {

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

	if IsExpired(m) {
		t.Fatalf("Macaroon has already expired")
	}
}
