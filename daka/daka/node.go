package daka

type Unimkt struct {
	Lotto Lotto `json:"lotto"`
}

type Province struct {
	Area AreaAccount `json:"area"`
}

type City struct {
	Area AreaAccount `json:"area"`
}

type Zone struct {
	Area AreaAccount `json:"area"`
}

type Branch struct {
	Community CommunityAccount `json:"community"`
}

type Division struct {
	Community CommunityAccount `json:"community"`
}

type Associate struct {
	Community CommunityAccount `json:"community"`
}

type Merchant struct {
	Merchant MerchantAccount `json:"merchant"`
	Lotto    LottoAccount    `json:"lotto"`
}

type Member struct {
	Member MemberAccount `json:"member"`
	Lotto  LottoAccount  `json:"lotto"`
}
