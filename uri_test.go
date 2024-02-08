package macaroon

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestNewMacaroonURI(t *testing.T) {

	expected := "macaroon://sfomuseum.org?duration=PT1M&key=file%3A%2F%2F%2Fusr%2Flocal%2Fsfomuseum%2Fgo-macaroon%2Ffixtures%2Fsigning.key"

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
		t.Fatalf("Failed to create new macaroon URI, %v", err)
	}

	if m_uri != expected {
		t.Fatalf("Expected '%s' but got '%s'", expected, m_uri)
	}
}
