package macaroon

import (
	"time"

	flyio_macaroon "github.com/superfly/macaroon"
)

func IsExpired(m *flyio_macaroon.Macaroon) bool {

	now := time.Now()
	expires := m.Expiration()

	if expires.After(now) {
		return false
	}

	return true
}
