package sfomuseum

import (
	"github.com/superfly/macaroon"
)

func init() {
	macaroon.RegisterCaveatType(&HasRoleCaveat{})
	macaroon.RegisterCaveatType(&IsUserCaveat{})
}
