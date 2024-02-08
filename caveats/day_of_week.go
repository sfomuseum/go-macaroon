package caveats

import (
	"fmt"
	"time"

	"github.com/superfly/macaroon"
)

func init() {
	macaroon.RegisterCaveatType(&DayOfWeekCaveat{})
}

type DayOfWeekCaveat struct {
	macaroon.Caveat
	Days []string `json:"days"`
}

func (c *DayOfWeekCaveat) CaveatType() macaroon.CaveatType {
	return macaroon.CavMinUserRegisterable + 3
}

func (c *DayOfWeekCaveat) Name() string {
	return "DayOfWeekCaveat"
}

func (c *DayOfWeekCaveat) Prohibits(f macaroon.Access) error {

	wf, isWF := f.(*DayOfWeekAccess)

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

type DayOfWeekAccess struct {
	macaroon.Access
	Day string
}

func (a *DayOfWeekAccess) Now() time.Time {
	return time.Now()
}

func (a *DayOfWeekAccess) Validate() error {
	return nil
}
