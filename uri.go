package macaroon

import (
	"net/url"
)

func NewMacaroonURI(loc string, key string, duration string) (string, error) {

	q := &url.Values{}
	q.Set("key", key)
	q.Set("duration", duration)

	u := &url.URL{}
	u.Scheme = "macaroon"
	u.Host = loc
	u.RawQuery = q.Encode()

	return u.String(), nil
}
