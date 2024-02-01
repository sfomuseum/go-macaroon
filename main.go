package main

// https://pkg.go.dev/github.com/superfly/macaroon
// https://fly.io/blog/macaroons-escalated-quickly/

import (
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/superfly/macaroon"
)

// Caveats

type IsUserCaveat struct {
	macaroon.Caveat
	Users []string `json:"users"`
}

func (c *IsUserCaveat) CaveatType() macaroon.CaveatType {
	return macaroon.CavMinUserRegisterable + 1
}

func (c *IsUserCaveat) Name() string {
	return "IsUserCaveat"
}

func (c *IsUserCaveat) Prohibits(f macaroon.Access) error {

	wf, isWF := f.(*IsUserAccess)

	if isWF {

		is_user := false

		for _, u := range c.Users {
			if u == wf.User {
				is_user = true
				break
			}
		}

		if !is_user {
			return fmt.Errorf("Invalid user")
		}
	}

	return nil
}

type HasRoleCaveat struct {
	macaroon.Caveat
	Role string `json:"role"`
}

func (c *HasRoleCaveat) CaveatType() macaroon.CaveatType {
	return macaroon.CavMinUserRegisterable + 2
}

func (c *HasRoleCaveat) Name() string {
	return "HasRoleCaveat"
}

func (c *HasRoleCaveat) Prohibits(f macaroon.Access) error {

	wf, isWF := f.(*HasRoleAccess)

	if isWF && wf.Role != c.Role {
		return fmt.Errorf("Invalid role")
	}

	return nil
}

// Accesses
// It is unclear to me what the purpose/requirement of the Validate method is...

type IsUserAccess struct {
	macaroon.Access
	User string
}

func (a *IsUserAccess) Now() time.Time {
	return time.Now()
}

func (a *IsUserAccess) Validate() error {
	return nil
}

type HasRoleAccess struct {
	macaroon.Access
	Role string
}

func (a *HasRoleAccess) Now() time.Time {
	return time.Now()
}

func (a *HasRoleAccess) Validate() error {
	return nil
}

func init() {
	macaroon.RegisterCaveatType(&HasRoleCaveat{})
	macaroon.RegisterCaveatType(&IsUserCaveat{})
}

func main() {

	kid := make([]byte, 0)
	loc := "sfomuseum.org"

	// Random every time
	key := macaroon.NewSigningKey()

	m, err := macaroon.New(kid, loc, key)

	if err != nil {
		log.Fatalf("Failed to create macaroon, %v", err)
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
		log.Fatalf("Failed to add caveat, %v", err)
	}

	b, err := m.Encode()

	if err != nil {
		log.Fatalf("Failed to encode, %v", err)
	}

	s := base64.StdEncoding.EncodeToString(b)

	b2, err := base64.StdEncoding.DecodeString(s)

	if err != nil {
		log.Fatalf("Failed to decode b2, %v", err)
	}

	log.Println("Expires", m.Expiration())
	log.Println("Sleeping...")
	time.Sleep(9 * time.Second)

	m2, err := macaroon.Decode(b2)

	if err != nil {
		log.Fatalf("Failed to decode, %v", err)
	}

	cs, err := m2.Verify(key, [][]byte{}, nil)

	if err != nil {
		log.Fatalf("Failed to verify, %v", err)
	}

	err = cs.Validate(
		&IsUserAccess{User: "alice"},
		&HasRoleAccess{Role: "staff"},
	)

	if err != nil {
		log.Fatalf("Failed to validate alice, %v", err)
	}

	log.Println("OK alice")

	err = cs.Validate(
		&IsUserAccess{User: "bob"},
		&HasRoleAccess{Role: "staff"},
	)

	if err != nil {
		log.Fatalf("Failed to validate bob, %v", err)
	}

	log.Println("OK bob")

	err = cs.Validate(
		&IsUserAccess{User: "doug"},
		&HasRoleAccess{Role: "staff"},
	)

	if err == nil {
		log.Fatalf("NO DOUG FOR YOU")
	}

	log.Println("DENY doug")

	log.Println("Sleep again...")
	time.Sleep(2 * time.Second)

	err = cs.Validate(
		&IsUserAccess{User: "bob"},
		&HasRoleAccess{Role: "staff"},
	)

	if err == nil {
		log.Fatalf("SHOULD HAVE EXPIRED")
	}

	log.Printf("Failed to validate bob, %v\n", err)
}
