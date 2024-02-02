package main

import (
	"log"

	"github.com/sfomuseum/go-macaroon/sfomuseum"
	"github.com/superfly/macaroon"
)

func main() {

	// Random every time
	key := macaroon.NewSigningKey()

	m, err := sfomuseum.TestMacaroon(key)

	if err != nil {
		log.Fatalf("Failed to create test macaroon, %v", err)
	}

	log.Println(m.String())

	ka := macaroon.NewEncryptionKey()

	c := &sfomuseum.DayOfWeekCaveat{
		Days: []string{"Friday"},
	}

	err = m.Add3P(ka, "example.com", c)

	if err != nil {
		log.Fatalf("Failed to add 3P, %v", err)
	}

	log.Println(m.String())
}
