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

	tp_loc := "example.com"

	err = m.Add3P(ka, tp_loc, c)

	if err != nil {
		log.Fatalf("Failed to add 3P, %v", err)
	}

	log.Println(m.String())

	enc, err := sfomuseum.EncodeMacaroonAsBase64(m)

	if err != nil {
		log.Fatalf("Failed to encode macaroon, %v", err)
	}

	m2, err := sfomuseum.DecodeMacaroonFromBase64(enc)

	if err != nil {
		log.Fatalf("Failed to decode macaroon, %v", err)
	}

	// 3P/discharge

	tp, err := m2.ThirdPartyTicket(tp_loc)

	if err != nil {
		log.Fatalf("Failed to get third party tickets, %v", err)
	}

	tp_caveats, tp_discharge, err := macaroon.DischargeTicket(ka, tp_loc, tp)

	if err != nil {
		log.Fatalf("Failed to parse discharge ticket, %v", err)
	}

	log.Println("TP", tp_caveats, tp_discharge)

	//

	cs, err := m2.Verify(key, [][]byte{}, nil)

	if err != nil {
		log.Fatalf("Failed to verify macaroon, %v", err)
	}

	for _, c := range cs.Caveats {
		log.Println(c.Name())
	}
}
