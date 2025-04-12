package nineora

import (
	"fmt"
	"github.com/hootuu/gelato/crtpto/md5x"
)

type NetworkID = NID

func NewNetworkID(chain Chain, addr Address) NetworkID {
	return md5x.MD5(fmt.Sprintf("[%s]%s", chain, addr))
}

type Network struct {
	NID       NetworkID `json:"nid"`
	NineoraID ID        `json:"nineora_id"`

	Chain   Chain   `json:"chain"`
	Nineora Address `json:"nineora"`

	Address   Address   `json:"address"`
	Symbol    string    `json:"symbol"`
	Owner     Address   `json:"owner"`
	Collar    Address   `json:"collar"`
	Ctrl      Ctrl      `json:"ctrl"`
	Committee Committee `json:"committee"`

	Fee uint64 `json:"fee"`

	TimestampMsUTC int64 `json:"timestamp_ms_utc"`
}
