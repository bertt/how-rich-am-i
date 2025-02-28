package model

type Coin struct {
	Id               string  `json:"id"`
	Name             string  `json:"name"`
	Symbol             string  `json:"symbol"`
	PriceUSD         float32 `json:"price_usd,string"`
	PriceEUR         float32 `json:"price_eur,string"`
	PriceBTC         float32 `json:"price_btc,string"`
	PercentChange1h  float32 `json:"percent_change_1h,string"`
	PercentChange24h float32 `json:"percent_change_24h,string"`
	PercentChange7d  float32 `json:"percent_change_7d,string"`
	LastUpdate       int     `json:"last_updated,string"`
}
