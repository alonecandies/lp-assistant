package models

import "encoding/json"

// Owner represents the owner field in the strategy response
// ...existing code...
type Owner struct {
	Address   string `json:"address"`
	Followers int    `json:"followers"`
}

type Protocol struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Logo    string `json:"logo"`
	Key     string `json:"key"`
}

type Token struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Decimals int    `json:"decimals"`
	Logo     string `json:"logo"`
	Tag      string `json:"tag"`
}

type StrategyProfile struct {
	ID                  int           `json:"id"`
	ChainID             int           `json:"chainId"`
	ChainName           string        `json:"chainName"`
	ChainLogo           string        `json:"chainLogo"`
	Owner               Owner         `json:"owner"`
	Protocol            Protocol      `json:"protocol"`
	Token0              Token         `json:"token0"`
	Token1              Token         `json:"token1"`
	LPValue             float64       `json:"lpValue"`
	Fees                float64       `json:"fees"`
	PNL                 float64       `json:"pnl"`
	ROI                 float64       `json:"roi"`
	APR                 float64       `json:"apr"`
	SubscribedValue     float64       `json:"subscribedValue"`
	RiskScore           float64       `json:"riskScore"`
	PerformanceChart    []interface{} `json:"performanceChart"`
	Status              string        `json:"status"`
	IsSupportAutomation bool          `json:"isSupportAutomation"`
	HasAutomationOrder  bool          `json:"hasAutomationOrder"`
	MinPrice            float64       `json:"minPrice"`
	MaxPrice            float64       `json:"maxPrice"`
	AgeInSeconds        int64         `json:"ageInSeconds"`
	CurrentPoolPrice    float64       `json:"currentPoolPrice"`
	PoolAddress         string        `json:"poolAddress"`
	InitialDepositValue float64       `json:"initialDepositValue"`
	VaultAddress        string        `json:"vaultAddress"`
}

// PoolDetail models

type PoolToken struct {
	Symbol   string      `json:"symbol"`
	Address  string      `json:"address"`
	Logo     string      `json:"logo"`
	Decimals json.Number `json:"decimals"`
	Balance  string      `json:"balance"`
}

type Tick struct {
	TickIdx        string `json:"tickIdx"`
	LiquidityGross string `json:"liquidityGross"`
	LiquidityNet   string `json:"liquidityNet"`
	Reserve0       string `json:"reserve0"`
	Reserve1       string `json:"reserve1"`
	Price0         string `json:"price0"`
	Price1         string `json:"price1"`
}

type PoolDetail struct {
	CurrentPrice0Usd float64   `json:"currentPrice0Usd"`
	CurrentPrice1Usd float64   `json:"currentPrice1Usd"`
	NfpmAddress      string    `json:"nfpmAddress"`
	TokenAddress0    string    `json:"tokenAddress0"`
	TokenAddress1    string    `json:"tokenAddress1"`
	Token0           PoolToken `json:"token0"`
	Token1           PoolToken `json:"token1"`
	CurrentPoolPrice float64   `json:"currentPoolPrice"`
	TVLUsd           string    `json:"tvlUsd"`
	VolumeUsd24h     string    `json:"volumeUsd24h"`
	FeeUsd24h        string    `json:"feeUsd24h"`
	Tag              string    `json:"tag"`
	Ticks            []Tick    `json:"ticks"`
	ChainID          int       `json:"chainId"`
	ChainLogo        string    `json:"chainLogo"`
	ProtocolName     string    `json:"protocolName"`
	ProtocolLogo     string    `json:"protocolLogo"`
	FeeTier          float64   `json:"feeTier"`
	TickSpacing      string    `json:"tickSpacing"`
}
