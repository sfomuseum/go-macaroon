package caveats

import (
	"fmt"
	"slices"
	"time"

	flyio_macaroon "github.com/superfly/macaroon"
)

func init() {
	flyio_macaroon.RegisterCaveatType(&EnsureAccountCaveat{})
}

type EnsureAccountCaveat struct {
	flyio_macaroon.Caveat
	AccountId int64    `json:"account_id"`
	RolesAny  []string `json:"roles_any,omitempty"`
	RolesAll  []string `json:"roles_all,omitempty"`
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

		if len(c.RolesAll) > 0 {

			if !slices.Equal(c.RolesAll, access.Roles) {
				return fmt.Errorf("Insufficient roles")
			}
		}

		if len(c.RolesAny) > 0 {

			has_role := false

			for _, r := range access.Roles {
				if slices.Contains(c.RolesAny, r) {
					has_role = true
					break
				}
			}

			if !has_role {
				return fmt.Errorf("Invalid role")
			}
		}
	}

	return nil
}

type EnsureAccountAccess struct {
	flyio_macaroon.Access
	AccountId int64
	Roles     []string
}

func (a *EnsureAccountAccess) Now() time.Time {
	return time.Now()
}

func (a *EnsureAccountAccess) Validate() error {
	return nil
}
