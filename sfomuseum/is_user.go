package sfomuseum

import (
	"fmt"
	"time"

	"github.com/superfly/macaroon"
)

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
