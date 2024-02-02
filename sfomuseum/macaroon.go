package sfomuseum

import (
	"fmt"
	"time"

	"github.com/superfly/macaroon"
)

func TestMacaroon(key macaroon.SigningKey) (*macaroon.Macaroon, error) {

	kid := make([]byte, 0)
	loc := "sfomuseum.org"

	m, err := macaroon.New(kid, loc, key)

	if err != nil {
		return nil, fmt.Errorf("Failed to create macaroon, %v", err)
	}

	// Max lifetime

	max_d := 10 * time.Second

	c := &macaroon.ValidityWindow{
		NotBefore: time.Now().Unix(),
		NotAfter:  time.Now().Add(max_d).Unix(),
	}

	// Allowed users

	c2 := &IsUserCaveat{
		Users: []string{"bob", "alice"},
	}

	// Required role

	c3 := &HasRoleCaveat{
		Role: "staff",
	}

	err = m.Add(c, c2, c3)

	if err != nil {
		return nil, fmt.Errorf("Failed to add caveat, %v", err)
	}

	return m, nil
}
