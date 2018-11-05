package bitmex

type Order struct {
	OrderID               string `json:",omitempty"`
	ClOrdID               string `json:",omitempty"`
	ClOrdLinkID           string `json:",omitempty"`
	Account               int    `json:",omitempty"`
	Symbol                string `json:",omitempty"`
	Side                  string `json:",omitempty"`
	SimpleOrderQty        int    `json:",omitempty"`
	OrderQty              int    `json:",omitempty"`
	Price                 int    `json:",omitempty"`
	DisplayQty            int    `json:",omitempty"`
	StopPx                int    `json:",omitempty"`
	PegOffsetValue        int    `json:",omitempty"`
	PegPriceType          string `json:",omitempty"`
	Currency              string `json:",omitempty"`
	SettlCurrency         string `json:",omitempty"`
	OrdType               string `json:",omitempty"`
	TimeInForce           string `json:",omitempty"`
	ExecInst              string `json:",omitempty"`
	ContingencyType       string `json:",omitempty"`
	ExDestination         string `json:",omitempty"`
	OrdStatus             string `json:",omitempty"`
	Triggered             string `json:",omitempty"`
	WorkingIndicator      bool   `json:",omitempty"`
	OrdRejReason          string `json:",omitempty"`
	SimpleLeavesQty       int    `json:",omitempty"`
	LeavesQty             int    `json:",omitempty"`
	SimpleCumQty          int    `json:",omitempty"`
	CumQty                int    `json:",omitempty"`
	AvgPx                 int    `json:",omitempty"`
	MultiLegReportingType string `json:",omitempty"`
	Text                  string `json:",omitempty"`
	TransactTime          string `json:",omitempty"`
	Timestamp             string `json:",omitempty"`
}
