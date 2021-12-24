package main

type orderVo struct {
	Time    string `json:"time"`
	TradeId int64  `json:"tradeId"`
	Price   string `json:"price"`
	Size    string `json:"size"`
	Side    string `json:"side"`
}
