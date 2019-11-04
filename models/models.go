package models

// Price .
type Price struct {
	Symbol   string  `json:"symbol"`
	PriceUSD float64 `json:"price_usd"`
}

// Toast .
type Toast struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
