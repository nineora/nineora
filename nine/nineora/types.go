package nineora

import (
	"fmt"
	"github.com/hootuu/gelato/crtpto/md5x"
	"github.com/hootuu/gelato/errors"
	"regexp"
)

type Link string

type LinkType string

func (lt LinkType) Assert(expected LinkType) *errors.Error {
	if lt != expected {
		return errors.Verify(fmt.Sprintf(
			"invalid link type[ expected: %s, but: %s ]",
			expected,
			lt,
		))
	}
	return nil
}

func (link Link) IsEmpty() bool {
	return link == ""
}

func (link Link) Verify() *errors.Error {
	if ok := md5x.IsMD5(string(link)); !ok {
		return errors.Verify("link invalid")
	}
	return nil
}

func LinkOf(network NetworkID, paras ...string) Link {
	return Link(md5x.New(network, ".", paras...))
}

type SeedType string

const (
	SeedEmail  = "EMAIL"
	SeedMobile = "MOBI"
)

type Seed struct {
	SeedType SeedType `json:"seed_type"`
	Seed     string   `json:"seed"`
}

var gEmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var gMobiRegex = regexp.MustCompile(`^\+[0-9]{7,15}$`)

func (seed Seed) Verify() *errors.Error {
	switch seed.SeedType {
	case SeedEmail:
		if ok := gEmailRegex.MatchString(seed.Seed); !ok {
			return errors.Verify("invalid seed email")
		}
		return nil
	case SeedMobile:
		if ok := gMobiRegex.MatchString(seed.Seed); !ok {
			return errors.Verify("invalid seed mobile phone number")
		}
		return nil
	default:
		return errors.Verify("seed type invalid")
	}
}

func SeedOfEmail(email string) Seed {
	return Seed{
		SeedType: SeedEmail,
		Seed:     email,
	}
}

func SeedOfCell(mobi string) Seed {
	return Seed{
		SeedType: SeedMobile,
		Seed:     mobi,
	}
}
