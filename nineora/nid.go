package nineora

import (
	"fmt"
	"github.com/hootuu/gelato/crtpto/md5x"
)

type NID = string

func NewNID(chain Chain, addr Address) NID {
	idStr := fmt.Sprintf("%s@%s", addr, chain)
	return md5x.MD5(idStr)
}
