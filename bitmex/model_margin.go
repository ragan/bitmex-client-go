package bitmex

type Margin struct {
	Account            int    `json:",omitempty"`
	Currency           string `json:",omitempty"`
	RiskLimit          int    `json:",omitempty"`
	PrevState          string `json:",omitempty"`
	State              string `json:",omitempty"`
	Action             string `json:",omitempty"`
	Amount             int    `json:",omitempty"`
	PendingCredit      int    `json:",omitempty"`
	PendingDebit       int    `json:",omitempty"`
	ConfirmedDebit     int    `json:",omitempty"`
	PrevRealisedPnl    int    `json:",omitempty"`
	PrevUnrealisedPnl  int    `json:",omitempty"`
	GrossComm          int    `json:",omitempty"`
	GrossOpenCost      int    `json:",omitempty"`
	GrossOpenPremium   int    `json:",omitempty"`
	GrossExecCost      int    `json:",omitempty"`
	GrossMarkValue     int    `json:",omitempty"`
	RiskValue          int    `json:",omitempty"`
	TaxableMargin      int    `json:",omitempty"`
	InitMargin         int    `json:",omitempty"`
	MaintMargin        int    `json:",omitempty"`
	SessionMargin      int    `json:",omitempty"`
	TargetExcessMargin int    `json:",omitempty"`
	VarMargin          int    `json:",omitempty"`
	RealisedPnl        int    `json:",omitempty"`
	UnrealisedPnl      int    `json:",omitempty"`
	IndicativeTax      int    `json:",omitempty"`
	UnrealisedProfit   int    `json:",omitempty"`
	SyntheticMargin    int    `json:",omitempty"`
	WalletBalance      int    `json:",omitempty"`
	MarginBalance      int    `json:",omitempty"`
	MarginBalancePcnt  int    `json:",omitempty"`
	MarginLeverage     int    `json:",omitempty"`
	MarginUsedPcnt     int    `json:",omitempty"`
	ExcessMargin       int    `json:",omitempty"`
	ExcessMarginPcnt   int    `json:",omitempty"`
	AvailableMargin    int    `json:",omitempty"`
	WithdrawableMargin int    `json:",omitempty"`
	Timestamp          string `json:",omitempty"`
	GrossLastValue     int    `json:",omitempty"`
	Commission         int    `json:",omitempty"`
}
