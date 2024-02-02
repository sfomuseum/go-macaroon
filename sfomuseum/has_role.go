package sfomuseum

import (
	"fmt"
	"time"

	"github.com/superfly/macaroon"
)

func init() {
	macaroon.RegisterCaveatType(&HasRoleCaveat{})
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
