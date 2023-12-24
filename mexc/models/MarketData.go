package models

// Check Server Time
// /api/v3/time
type ServerTimeResponse struct {
	ServerTime int64 `json:"serverTime"`
}

// API default symbol
// /api/v3/defaultSymbols
type DefaultSymbolsResponse struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
	Msg  string   `json:"msg"`
}

// Exchange Information
// /api/v3/exchangeInfo
type ExchangeInfoRequest struct {
	Timezone                   string   `json:"timezone"`
	ServerTime                 int64    `json:"serverTime"`
	RateLimits                 []int    `json:"rateLimits"`
	ExchangeFilters            []string `json:"exchangeFilters"`
	Symbol                     string   `json:"symbol"`
	Status                     string   `json:"status"`
	BaseAsset                  string   `json:"baseAsset"`
	BaseAssetPrecision         int      `json:"baseAssetPrecision"`
	QuoteAsset                 string   `json:"quoteAsset"`
	QuotePrecision             int      `json:"quotePrecision"`
	QuoteAssetPrecision        int      `json:"quoteAssetPrecision"`
	BaseCommissionPrecision    int      `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision   int      `json:"quoteCommissionPrecision"`
	OrderTypes                 []string `json:"orderTypes"`
	QuoteOrderQtyMarketAllowed bool     `json:"quoteOrderQtyMarketAllowed"`
	IsSpotTradingAllowed       bool     `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed     bool     `json:"isMarginTradingAllowed"`
	Permissions                []string `json:"permissions"`
	MaxQuoteAmount             string   `json:"maxQuoteAmount"`
	MakerCommission            string   `json:"makerCommission"`
	TakerCommission            string   `json:"takerCommission"`
	QuoteAmountPrecision       string   `json:"quoteAmountPrecision"`
	BaseSizePrecision          string   `json:"baseSizePrecision"`
	QuoteAmountPrecisionMarket string   `json:"quoteAmountPrecisionMarket"`
	MaxQuoteAmountMarket       string   `json:"maxQuoteAmountMarket"`
}

// Order Book
// /api/v3/depth
type Bid struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

type Ask struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

type OrderBookResponse struct {
	LastUpdateId int   `json:"lastUpdateId"`
	Bids         []Bid `json:"bids"`
	Asks         []Ask `json:"asks"`
}

// Recent Trades List
// /api/v3/trades
type RecentTradesListResponse struct {
	Id           string `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int    `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

// Compressed/Aggregate Trades List
// /api/v3/aggTrades
type AggregateTradesResponse struct {
	AggregateTradeId string `json:"a"`
	FirstTradeId     string `json:"f"`
	LastTradeId      string `json:"l"`
	Price            string `json:"p"`
	Quantity         string `json:"q"`
	Timestamp        int    `json:"t"`
	IsBuyerMaker     bool   `json:"m"`
	IsBestMatch      bool   `json:"M"`
}

// Kline/Candlestick Data
// /api/v3/klines
type KlineDataResponse struct {
	OpenTime         int
	Open             string
	High             string
	Low              string
	Close            string
	Volume           string
	CloseTime        int
	QuoteAssetVolume string
}

// Current Average Price
// /api/v3/avgPrice

type CurrentAveragePriceResponse struct {
	Mins  int    `json:"mins"`
	Price string `json:"price"`
}

// 24hr Ticker Price Change Statistics
// /api/v3/ticker/24hr

type TickerPriceChangeStatisticsResponse struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int    `json:"openTime"`
	CloseTime          int    `json:"closeTime"`
	Count              string `json:"count"`
}

// Symbol Price Ticker
///api/v3/ticker/price

type SymbolPriceTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type SymbolPricesTicker struct {
	_ []SymbolPriceTicker
}

type SymbolPriceTickerResponse struct {
	_ SymbolPriceTicker
}

type SymbolPricesTickerResponse struct {
	_ SymbolPricesTicker
}

// Symbol Order Book Ticker
// /api/v3/ticker/bookTicker
type SymbolOrderBookTicker struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

type SymbolOrderBookTickers struct {
	_ []SymbolOrderBookTicker
}

type SymbolOrderBookTickerResponse struct {
	_ SymbolOrderBookTicker
}

type SymbolOrderBookTickersResponse struct {
	_ SymbolOrderBookTickers
}
