package caveats

import (
	"fmt"
	"time"

	"github.com/superfly/macaroon"
)

func init() {
	macaroon.RegisterCaveatType(&isUserCaveat{})
	macaroon.RegisterCaveatType(&hasRoleCaveat{})
	macaroon.RegisterCaveatType(&dayOfWeekCaveat{})
}

type isUserCaveat struct {
	macaroon.Caveat
	Users []string `json:"users"`
}

func (c *isUserCaveat) CaveatType() macaroon.CaveatType {
	return CavSFOMuseumTestIsUser
}

func (c *isUserCaveat) Name() string {
	return "isUserCaveat"
}

func (c *isUserCaveat) Prohibits(f macaroon.Access) error {

	wf, isWF := f.(*isUserAccess)

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

type isUserAccess struct {
	macaroon.Access
	User string
}

func (a *isUserAccess) Now() time.Time {
	return time.Now()
}

func (a *isUserAccess) Validate() error {
	return nil
}

type hasRoleCaveat struct {
	macaroon.Caveat
	Role string `json:"role"`
}

func (c *hasRoleCaveat) CaveatType() macaroon.CaveatType {
	return CavSFOMuseumTestHasRole
}

func (c *hasRoleCaveat) Name() string {
	return "hasRoleCaveat"
}

func (c *hasRoleCaveat) Prohibits(f macaroon.Access) error {

	wf, isWF := f.(*hasRoleAccess)

	if isWF && wf.Role != c.Role {
		return fmt.Errorf("Invalid role")
	}

	return nil
}

type hasRoleAccess struct {
	macaroon.Access
	Role string
}

func (a *hasRoleAccess) Now() time.Time {
	return time.Now()
}

func (a *hasRoleAccess) Validate() error {
	return nil
}

type dayOfWeekCaveat struct {
	macaroon.Caveat
	Days []string `json:"days"`
}

func (c *dayOfWeekCaveat) CaveatType() macaroon.CaveatType {
	return CavSFOMuseumTestDayOfWeek
}

func (c *dayOfWeekCaveat) Name() string {
	return "dayOfWeekCaveat"
}

func (c *dayOfWeekCaveat) Prohibits(f macaroon.Access) error {

	wf, isWF := f.(*dayOfWeekAccess)

	if isWF {

		is_dow := false

		for _, d := range c.Days {
			if d == wf.Day {
				is_dow = true
				break
			}
		}

		if !is_dow {
			return fmt.Errorf("Invalid day")
		}
	}

	return nil
}

type dayOfWeekAccess struct {
	macaroon.Access
	Day string
}

func (a *dayOfWeekAccess) Now() time.Time {
	return time.Now()
}

func (a *dayOfWeekAccess) Validate() error {
	return nil
}
