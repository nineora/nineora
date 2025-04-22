package daka

const (
	RoundDay     Round = 1
	RoundWeek    Round = 2
	RoundMonth   Round = 3
	RoundQuarter Round = 4
	RoundYear    Round = 5
)

type Round int

func (r Round) ToString() string {
	switch r {
	case RoundDay:
		return "day"
	case RoundWeek:
		return "week"
	case RoundMonth:
		return "month"
	case RoundQuarter:
		return "quarter"
	case RoundYear:
		return "year"
	}
	return ""
}

type Lotto struct {
	EcologyRate uint64 `json:"ecology_rate"`
}

type Pool struct {
	Score        uint64 `json:"score"`
	Bonus        uint64 `json:"bonus"`
	BonusBalance uint64 `json:"bonus_balance"`
}

type Item struct {
	Numb   uint64 `json:"numb"`
	Open   uint64 `json:"open"`
	Profit uint64 `json:"profit"`
	Cur    uint64 `json:"cur"`
	Nxt    uint64 `json:"nxt"`
}

type Activity struct {
	Weight      uint64 `json:"weight"`
	Amount      uint64 `json:"amount"`
	AmountMonth uint64 `json:"amount_month"`
	Times       uint64 `json:"times"`
	TimesMonth  uint64 `json:"times_month"`
	LstTime     uint64 `json:"lst_time"`
	IniTime     uint64 `json:"ini_time"`
}

type LottoAccount struct {
	InvestRatio uint64   `json:"invest_ratio"`
	ProfitRatio uint64   `json:"profit_ratio"`
	Day         Item     `json:"day"`
	Week        Item     `json:"week"`
	Month       Item     `json:"month"`
	Quarter     Item     `json:"quarter"`
	Year        Item     `json:"year"`
	Activity    Activity `json:"activity"`
}
