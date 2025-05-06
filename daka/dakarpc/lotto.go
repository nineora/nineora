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
	PublicSrcMerchant  nineora.Link `json:"public_src_merchant"`
}

type LottoMemberClaimReq struct {
	Member nineora.Link `json:"member"`
}

type LottoMerchantClaimReq struct {
	Merchant nineora.Link `json:"merchant"`
}

type LottoAreaSetReq struct {
	ProvinceRatio uint64 `json:"province_ratio"`
	CityRatio     uint64 `json:"city_ratio"`
	ZoneRatio     uint64 `json:"zone_ratio"`
}

type LottoCommunitySetReq struct {
	BranchRatio    uint64 `json:"branch_ratio"`
	DivisionRatio  uint64 `json:"division_ratio"`
	AssociateRatio uint64 `json:"associate_ratio"`
}

type LottoSupplySetReq struct {
	ShiftRatio     uint64 `json:"shift_ratio"`
	LockedRatio    uint64 `json:"locked_ratio"`
	NxtUnlockRatio uint64 `json:"nxt_unlock_ratio"`
}

const (
	LottoPath              = "daka/lotto"
	LottoTriggerPath       = LottoPath + "/trigger"
	LottoMemberClaimPath   = LottoPath + "/claim/member"
	LottoMerchantClaimPath = LottoPath + "/claim/merchant"
	LottoAreaSetPath       = LottoPath + "/set/area"
	LottoCommunitySetPath  = LottoPath + "/set/community"
	LottoSupplySetPath     = LottoPath + "/set/supply"
)

type LottoTrigger func(req *LottoTriggerReq) (*chain.Tx, *errors.Error)
type LottoMemberClaim func(req *LottoMemberClaimReq) (*chain.Tx, *errors.Error)
type LottoMerchantClaim func(req *LottoMerchantClaimReq) (*chain.Tx, *errors.Error)
type LottoAreaSet func(req *LottoAreaSetReq) (*chain.Tx, *errors.Error)
type LottoCommunitySet func(req *LottoCommunitySetReq) (*chain.Tx, *errors.Error)
type LottoSupplySet func(req *LottoSupplySetReq) (*chain.Tx, *errors.Error)
