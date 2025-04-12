package nineora

type Event struct {
	Chain          Chain                  `json:"chain"`
	EventType      string                 `json:"event_type"`
	CtxType        string                 `json:"ctx_type"`
	Biz            uint16                 `json:"biz"`
	Memo           string                 `json:"memo"`
	Data           map[string]interface{} `json:"data"`
	TimestampMsUTC int64                  `json:"timestamp_ms_utc"`
}
