package caveats

import (
	"fmt"
	"time"

	flyio_macaroon "github.com/superfly/macaroon"
)

func init() {
	flyio_macaroon.RegisterCaveatType(&EnsureAccountCaveat{})
}

type EnsureAccountCaveat struct {
	flyio_macaroon.Caveat
	AccountId int64    `json:"account_id"`
	Roles     []string `json:"roles,omitempty"`
}

func (c *EnsureAccountCaveat) CaveatType() flyio_macaroon.CaveatType {
	return CavSFOMuseumEnsureAccount
}

func (c *EnsureAccountCaveat) Name() string {
	return "EnsureAccountCaveat"
}

func (c *EnsureAccountCaveat) Prohibits(f flyio_macaroon.Access) error {

	access, is_access := f.(*EnsureAccountAccess)

	if is_access {

		if access.AccountId != c.AccountId {
			return fmt.Errorf("Invalid account")
		}
	}

	return nil
}

type EnsureAccountAccess struct {
	flyio_macaroon.Access
	AccountId int64
}

func (a *EnsureAccountAccess) Now() time.Time {
	return time.Now()
}

func (a *EnsureAccountAccess) Validate() error {
	return nil
}
