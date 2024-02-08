package macaroon

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
)

func TestNewEncryptionKey(t *testing.T) {

	ctx := context.Background()

	abs_path, err := filepath.Abs(".")

	if err != nil {
		t.Fatalf("Failed to derive absolute path, %v", err)
	}

	path_key := filepath.Join(abs_path, "fixtures/3p-encryption.key")
	key_uri := fmt.Sprintf("file://%s", path_key)

	_, err = NewEncryptionKey(ctx, key_uri)

	if err != nil {
		t.Fatalf("Failed to derive encryption key from %s, %v", key_uri, err)
	}
}
