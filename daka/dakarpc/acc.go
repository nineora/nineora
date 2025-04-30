package dakarpc

import (
	"github.com/hootuu/gelato/errors"
	"github.com/nineora/nineora/daka/daka"
	"github.com/nineora/nineora/nine/nineora"
)

const (
	ReadPath         = "daka/read"
	ProvinceGetPath  = ReadPath + "/province"
	CityGetPath      = ReadPath + "/city"
	ZoneGetPath      = ReadPath + "/zone"
	BranchGetPath    = ReadPath + "/branch"
	DivisionGetPath  = ReadPath + "/division"
	AssociateGetPath = ReadPath + "/associate"
	MerchantGetPath  = ReadPath + "/merchant"
	MemberGetPath    = ReadPath + "/member"
	LottoGetPath     = ReadPath + "/lotto/info"
)

type ProvinceGet func(link nineora.Link) (*daka.Province, *errors.Error)
type CityGet func(link nineora.Link) (*daka.City, *errors.Error)
type ZoneGet func(link nineora.Link) (*daka.Zone, *errors.Error)

type BranchGet func(link nineora.Link) (*daka.Branch, *errors.Error)
type DivisionGet func(link nineora.Link) (*daka.Division, *errors.Error)
type AssociateGet func(link nineora.Link) (*daka.Associate, *errors.Error)

type MerchantGet func(link nineora.Link) (*daka.Merchant, *errors.Error)

type MemberGet func(link nineora.Link) (*daka.Member, *errors.Error)

type LottoGet func() (*daka.Lotto, *errors.Error)
