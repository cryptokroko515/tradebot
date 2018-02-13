package entity

import "time"

type UserEntity interface {
	GetId() uint
	GetUsername() string
	GetLocalCurrency() string
	GetEtherbase() string
	GetKeystore() string
	GetWallets() []UserWallet
	GetExchanges() []UserCryptoExchange
}

type UserWalletEntity interface {
	GetUserId() uint
	GetCurrency() string
	GetAddress() string
}

type UserExchangeEntity interface {
	GetUserId() uint
	GetName() string
	GetURL() string
	GetKey() string
	GetSecret() string
	GetExtra() string
}

type ChartEntity interface {
	GetId() uint
	GetUserId() uint
	GetBase() string
	GetQuote() string
	GetPeriod() int
	GetExchangeName() string
	IsAutoTrade() bool
	GetAutoTrade() uint
	SetIndicators(indicators []ChartIndicator)
	GetIndicators() []ChartIndicator
	AddIndicator(indicator *ChartIndicator)
	SetStrategies(strategies []ChartStrategy)
	GetStrategies() []ChartStrategy
	AddStrategy(strategy *ChartStrategy)
	SetTrades(trades []Trade)
	GetTrades() []Trade
	AddTrade(trade Trade)
}

type ChartIndicatorEntity interface {
	GetId() uint
	GetChartId() uint
	GetName() string
	GetParameters() string
}

type ChartStrategyEntity interface {
	GetId() uint
	GetChartId() uint
	GetName() string
	GetParameters() string
}

type TradeEntity interface {
	GetId() uint
	GetChartId() uint
	GetUserId() uint
	GetBase() string
	GetQuote() string
	GetExchangeName() string
	GetDate() time.Time
	GetType() string
	GetPrice() float64
	GetAmount() float64
	GetChartData() string
}