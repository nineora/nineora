package dakarpc

import (
	"github.com/hootuu/gelato/errors"
	"github.com/nineora/nineora/nine/chain"
	"github.com/nineora/nineora/nine/nineora"
)

type LottoTriggerReq struct {
	OrderAmount        uint64       `json:"order_amount"`
	ContributionAmount uint64       `json:"contribution_amount"`
	Member             nineora.Link `json:"member"`
	Merchant           nineora.Link `json:"merchant"`
	Zone               nineora.Link `json:"zone"`
}

type LottoMemberClaimReq struct {
	Member nineora.Link `json:"member"`
}

type LottoMerchantClaimReq struct {
	Merchant nineora.Link `json:"merchant"`
}

const (
	LottoPath              = "daka/lotto"
	LottoTriggerPath       = LottoPath + "/trigger"
	LottoMemberClaimPath   = LottoPath + "/claim/member"
	LottoMerchantClaimPath = LottoPath + "/claim/merchant"
)

type LottoTrigger func(req *LottoTriggerReq) (*chain.Tx, *errors.Error)
type LottoMemberClaim func(req *LottoMemberClaimReq) (*chain.Tx, *errors.Error)
type LottoMerchantClaim func(req *LottoMerchantClaimReq) (*chain.Tx, *errors.Error)
