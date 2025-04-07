package ninerpc

import (
	"github.com/hootuu/gelato/errors"
	"github.com/nineora/nineora/nineora"
)

type IDomain interface {
	InitDomain(nineoraID nineora.ID, symbol string) (*nineora.Domain, *errors.Error)
}
