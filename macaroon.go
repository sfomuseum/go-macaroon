package macaroon

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/sfomuseum/iso8601duration"
	flyio_macaroon "github.com/superfly/macaroon"
)

type NewMacaroonOptions struct {
	SigningKey flyio_macaroon.SigningKey
	Duration   time.Duration
	Location   string
	Body       []byte
}

func NewMacaroon(ctx context.Context, uri string) (*flyio_macaroon.Macaroon, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse macaroon URI, %w", err)
	}

	q := u.Query()

	key_loc := u.Host
	key_uri := q.Get("key")
	key_duration := q.Get("duration")

	k, err := NewSigningKey(ctx, key_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive signing key from ?key= parameter, %w", err)
	}

	d, err := duration.FromString(key_duration)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse ?duration= parameter, %w", err)
	}

	opts := &NewMacaroonOptions{
		SigningKey: k,
		Duration:   d.ToDuration(),
		Location:   key_loc,
	}

	return NewMacaroonWithOptions(ctx, opts)
}

func NewMacaroonWithOptions(ctx context.Context, opts *NewMacaroonOptions) (*flyio_macaroon.Macaroon, error) {

	body := opts.Body // make([]byte, 0)

	m, err := flyio_macaroon.New(body, opts.Location, opts.SigningKey)

	if err != nil {
		return nil, fmt.Errorf("Failed to create macaroon, %v", err)
	}

	c := &flyio_macaroon.ValidityWindow{
		NotBefore: time.Now().Unix(),
		NotAfter:  time.Now().Add(opts.Duration).Unix(),
	}

	err = m.Add(c)

	if err != nil {
		return nil, fmt.Errorf("Failed to add caveat, %v", err)
	}

	return m, nil
}
