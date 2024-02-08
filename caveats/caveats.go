package caveats

import (
	flyio_macaroon "github.com/superfly/macaroon"
)

const (
	CavSFOMuseumTestIsUser = flyio_macaroon.CavMinUserRegisterable << iota
	CavSFOMuseumTestHasRole
	CavSFOMuseumTestDayOfWeek

	CavSFOMuseumEnsureAccount
)
