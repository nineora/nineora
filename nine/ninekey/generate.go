package ninekey

import "github.com/hootuu/gelato/errors"

type Generate func() (*Key, *errors.Error)

var gGenerateFactory map[Type]Generate

func init() {
	gGenerateFactory = make(map[Type]Generate)
	gGenerateFactory[Ed25519] = Ed25519Generate
}

func Gen(algoType Type) (*Key, *errors.Error) {
	genFunc, ok := gGenerateFactory[algoType]
	if !ok {
		return nil, errors.System("can not support this type: " + algoType.String())
	}
	return genFunc()
}
