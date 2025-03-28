package macaroon

import (
	"fmt"
	"net/url"
	"path/filepath"
	"testing"
)

func TestNewMacaroonURI(t *testing.T) {

	abs_path, err := filepath.Abs("fixtures/signing.key")

	if err != nil {
		t.Fatalf("Failed to derive absolute path for signing key, %v", err)
	}

	key_uri := fmt.Sprintf("file://%s", abs_path)
	enc_uri := url.QueryEscape(key_uri)

	loc := "sfomuseum.org"
	duration := "PT1M"

	expected := fmt.Sprintf("macaroon://%s?duration=%s&key=%s", loc, duration, enc_uri)

	m_uri, err := NewMacaroonURI(loc, key_uri, duration)

	if err != nil {
		t.Fatalf("Failed to create new macaroon URI, %v", err)
	}

	if m_uri != expected {
		t.Fatalf("Expected '%s' but got '%s'", expected, m_uri)
	}
}
