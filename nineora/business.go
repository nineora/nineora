package nineora

import (
	"fmt"
	"github.com/hootuu/gelato/crtpto/md5x"
)

type BusinessID = NID

func NewBusinessID(chain Chain, addr Address) BusinessID {
	return md5x.MD5(fmt.Sprintf("[%s]%s", chain, addr))
}

type Business struct {
	NID       BusinessID `json:"nid"`
	NineoraID ID         `json:"nineora_id"`
	NetworkID NetworkID  `json:"network_id"`

	Chain     Chain     `json:"chain"`
	Nineora   Address   `json:"nineora"`
	Network   Address   `json:"domain"`
	Address   Address   `json:"address"`
	Symbol    string    `json:"symbol"`
	Owner     Address   `json:"owner"`
	Collar    Address   `json:"collar"`
	Ctrl      Ctrl      `json:"ctrl"`
	Committee Committee `json:"committee"`

	BizType string                 `json:"biz_type"`
	Biz     map[string]interface{} `json:"biz"`

	TimestampMsUTC int64 `json:"timestamp_ms_utc"`
}

type Biz[T any] struct {
	Biz T `json:"biz"`
}
