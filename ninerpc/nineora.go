package ninerpc

import (
	"github.com/hootuu/gelato/errors"
	"github.com/nineora/nineora/nineora"
)

type NineoraGetReq struct {
}

const (
	NineoraPath    = "nineora"
	NineoraGetPath = NineoraPath + "/get"
)

type NineoraGet func(req *NineoraGetReq) (*nineora.Nineora, *errors.Error)
