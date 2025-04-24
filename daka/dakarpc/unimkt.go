package dakarpc

import (
	"github.com/hootuu/gelato/errors"
	"github.com/nineora/nineora/nine/chain"
	"github.com/nineora/nineora/nine/nineora"
)

type ProvinceCreateReq struct {
	Link    nineora.Link `json:"link"`
	Benefit nineora.Link `json:"benefit"` //Benefit Account Link
}

type CityCreateReq struct {
	Link     nineora.Link `json:"link"`
	Province nineora.Link `json:"province"`
	Benefit  nineora.Link `json:"benefit"` //Benefit Account Link
}

type ZoneCreateReq struct {
	Link    nineora.Link `json:"link"`
	City    nineora.Link `json:"city"`
	Benefit nineora.Link `json:"benefit"` //Benefit Account Link
}

type BranchCreateReq struct {
	Link    nineora.Link `json:"link"`
	Benefit nineora.Link `json:"benefit"` //Benefit Account Link
}

type DivisionCreateReq struct {
	Link    nineora.Link `json:"link"`
	Branch  nineora.Link `json:"branch"`
	Benefit nineora.Link `json:"benefit"` //Benefit Account Link
}

type AssociateCreateReq struct {
	Link     nineora.Link `json:"link"`
	Division nineora.Link `json:"division"`
	Benefit  nineora.Link `json:"benefit"` //Benefit Account Link
}

type MerchantCreateReq struct {
	Link        nineora.Link `json:"link"`
	Associate   nineora.Link `json:"associate"`
	Benefit     nineora.Link `json:"benefit"`      //Benefit Account Link
	InvestRatio uint64       `json:"invest_ratio"` //Base 10000
}

type MemberCreateReq struct {
	Link     nineora.Link `json:"link"`
	Merchant nineora.Link `json:"merchant"`
	Benefit  nineora.Link `json:"benefit"` //Benefit Account Link
}

type MemberWithdrawReq struct {
	Member nineora.Link `json:"member"`
}

type MerchantWithdrawReq struct {
	Merchant nineora.Link `json:"merchant"`
}

type AssociateWithdrawReq struct {
	Associate nineora.Link `json:"associate"`
}

type DivisionWithdrawReq struct {
	Division nineora.Link `json:"division"`
}

type BranchWithdrawReq struct {
	Branch nineora.Link `json:"branch"`
}

type ZoneWithdrawReq struct {
	Zone nineora.Link `json:"zone"`
}

type CityWithdrawReq struct {
	City nineora.Link `json:"city"`
}

type ProvinceWithdrawReq struct {
	Province nineora.Link `json:"province"`
}

const (
	UnimktPath            = "daka/unimkt"
	ProvinceCreatePath    = UnimktPath + "/province/create"
	ProvinceWithdrawPath  = UnimktPath + "/province/withdraw"
	CityCreatePath        = UnimktPath + "/city/created"
	CityWithdrawPath      = UnimktPath + "/city/withdraw"
	ZoneCreatePath        = UnimktPath + "/zone/created"
	ZoneWithdrawPath      = UnimktPath + "/zone/withdraw"
	BranchCreatePath      = UnimktPath + "/branch/created"
	BranchWithdrawPath    = UnimktPath + "/branch/withdraw"
	DivisionCreatePath    = UnimktPath + "/division/created"
	DivisionWithdrawPath  = UnimktPath + "/division/withdraw"
	AssociateCreatePath   = UnimktPath + "/associate/created"
	AssociateWithdrawPath = UnimktPath + "/associate/withdraw"
	MerchantCreatePath    = UnimktPath + "/merchant/created"
	MerchantWithdrawPath  = UnimktPath + "/merchant/withdraw"
	MemberCreatePath      = UnimktPath + "/member/created"
	MemberWithdrawPath    = UnimktPath + "/member/withdraw"
)

type ProvinceCreate func(req *ProvinceCreateReq) (*chain.Tx, *errors.Error)
type CityCreate func(req *CityCreateReq) (*chain.Tx, *errors.Error)
type ZoneCreate func(req *ZoneCreateReq) (*chain.Tx, *errors.Error)
type BranchCreate func(req *BranchCreateReq) (*chain.Tx, *errors.Error)
type DivisionCreate func(req *DivisionCreateReq) (*chain.Tx, *errors.Error)
type AssociateCreate func(req *AssociateCreateReq) (*chain.Tx, *errors.Error)
type MerchantCreate func(req *MerchantCreateReq) (*chain.Tx, *errors.Error)
type MemberCreate func(req *MemberCreateReq) (*chain.Tx, *errors.Error)
type MemberWithdraw func(req *MemberWithdrawReq) (*chain.Tx, *errors.Error)
type MerchantWithdraw func(req *MerchantWithdrawReq) (*chain.Tx, *errors.Error)
type ProvinceWithdraw func(req *ProvinceWithdrawReq) (*chain.Tx, *errors.Error)
type CityWithdraw func(req *CityWithdrawReq) (*chain.Tx, *errors.Error)
type ZoneWithdraw func(req *ZoneWithdrawReq) (*chain.Tx, *errors.Error)
type BranchWithdraw func(req *BranchWithdrawReq) (*chain.Tx, *errors.Error)
type DivisionWithdraw func(req *DivisionWithdrawReq) (*chain.Tx, *errors.Error)
type AssociateWithdraw func(req *AssociateWithdrawReq) (*chain.Tx, *errors.Error)
