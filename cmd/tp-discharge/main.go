package main

// https://github.com/superfly/macaroon/blob/main/macaroon-thought.md#how-third-party-caveats-work
// https://github.com/superfly/macaroon/blob/main/tp/README.md

import (
	"log"
	"time"

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

	// START OF add third-party caveat here

	ka := macaroon.NewEncryptionKey()

	tp_c := &sfomuseum.DayOfWeekCaveat{
		Days: []string{"Friday"},
	}

	tp_loc := "example.com"

	err = m.Add3P(ka, tp_loc, tp_c)

	if err != nil {
		log.Fatalf("Failed to add 3P, %v", err)
	}

	// END OF add third-party caveat here

	enc, err := sfomuseum.EncodeMacaroonAsBase64(m)

	if err != nil {
		log.Fatalf("Failed to encode macaroon, %v", err)
	}

	// Pretend we are sending enc somewhere...

	log.Println("Sleeping 2 seconds to simulate exchanging macaroon")
	time.Sleep(2 * time.Second)

	// Pretend enc is being received somewhere...

	m2, err := sfomuseum.DecodeMacaroonFromBase64(enc)

	if err != nil {
		log.Fatalf("Failed to decode macaroon, %v", err)
	}

	// START OF third-party discharges

	discharges := make([][]byte, 0)
	discharge_keys := make(map[string]macaroon.EncryptionKey)

	tp, err := m2.ThirdPartyTicket(tp_loc)

	if err != nil {
		log.Fatalf("Failed to get third party tickets, %v", err)
	}

	// Pretend we are sending tp to loc here...

	log.Println("Sleeping 2 seconds to simulate sending discharge ticket")
	time.Sleep(2 * time.Second)

	tp_caveats, tp_discharge, err := macaroon.DischargeTicket(ka, tp_loc, tp)

	if err != nil {
		log.Fatalf("Failed to parse discharge ticket, %v", err)
	}

	tp_cs := macaroon.NewCaveatSet(tp_caveats...)

	err = tp_cs.Validate(
		&sfomuseum.DayOfWeekAccess{Day: "Friday"},
	)

	if err != nil {
		// Pretend tp_loc is sending back an error here
		log.Fatalf("TP dispatch failed to validate CS, %v", err)
	}

	encd, err := tp_discharge.Encode()

	if err != nil {
		log.Fatalf("Failed to create discharge token, %v", err)
	}

	// Pretend tp_loc is sending back 'encd' here...

	log.Println("Sleeping 2 seconds to simulate receiving discharge token")
	time.Sleep(2 * time.Second)

	discharges = append(discharges, encd)
	discharge_keys[tp_loc] = ka

	// END OF third-party discharges

	// Verify with all the other caveats

	cs, err := m2.Verify(key, discharges, discharge_keys)

	if err != nil {
		log.Fatalf("Failed to verify macaroon, %v", err)
	}

	// Test caveats (defined in sfomuseum.TestMacaroon)

	err = cs.Validate(
		&sfomuseum.IsUserAccess{User: "alice"},
		&sfomuseum.HasRoleAccess{Role: "staff"},
	)

	if err != nil {
		log.Fatalf("Failed to validate alice, %v", err)
	}

	log.Println("All caveats validate")

	log.Println("Sleeping 4 seconds to trigger expiration")
	time.Sleep(4 * time.Second)

	err = cs.Validate(
		&sfomuseum.IsUserAccess{User: "alice"},
		&sfomuseum.HasRoleAccess{Role: "staff"},
	)

	if err == nil {
		log.Fatalf("Expected macaroon to be expired")
	}

	log.Printf("Macaroon expired, %v", err)
}
