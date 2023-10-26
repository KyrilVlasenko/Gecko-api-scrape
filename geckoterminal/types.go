package geckoterminal

import "time"

type basicStruct struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// NetworksItem
type NetworksItem struct {
	basicStruct
	Attributes AttributesItem `json:"attributes"`
}

// AttributesItem
type AttributesItem struct {
	Name                     string  `json:"name"`
	CoingeckoAssetPlatformID *string `json:"coingecko_asset_platform_id,omitempty"`
}

// LinksItem used for pagination
type LinksItem struct {
	First string  `json:"first"`
	Prev  *string `json:"prev"`
	Next  *string `json:"next"`
	Last  string  `json:"last"`
}

// PoolDataItem
type PoolDataItem struct {
	basicStruct
	Attributes    PoolAttributesItem     `json:"attributes"`
	Relationships map[string]basicStruct `json:"relationships"`
}

// PoolAttributesItem
type PoolAttributesItem struct {
	BaseTokenPriceUSD             *string                     `json:"base_token_price_usd"`
	BaseTokenPriceNativeCurrency  *string                     `json:"base_token_price_native_currency"`
	QuotaTokenPriceUSD            *string                     `json:"quota_token_price_usd"`
	QuoteTokenPriceNativeCurrency *string                     `json:"quote_token_price_native_currency"`
	BaseTokenPriceQuoteToken      *string                     `json:"base_token_price_quote_token"`
	QuoteTokenPriceBaseToken      *string                     `json:"quote_token_price_base_token"`
	Address                       string                      `json:"address"`
	Name                          string                      `json:"name"`
	PoolCreatedAt                 *time.Time                  `json:"pool_created_at"`
	TokenPriceUSD                 *string                     `json:"token_price_usd,omitempty"`
	FDVUsed                       *string                     `json:"fdv_used"`
	MarketCapUSD                  *string                     `json:"market_cap_usd"`
	PriceChangePercentage         map[string]string           `json:"price_change_percentage"`
	Transactions                  map[string]TransactionsItem `json:"transactions"`
	VolumeUSD                     map[string]string           `json:"volume_usd"`
	ReserveInUSD                  *string                     `json:"reserve_in_usd"`
}

// TransactionsItem
type TransactionsItem struct {
	Buys  int64 `json:"buys"`
	Sells int64 `json:"sells"`
}

// PoolIncludedItem
type PoolIncludedItem struct {
	basicStruct
	Attributes struct {
		Address         string  `json:"address,omitempty"`
		Name            string  `json:"name"`
		Symbol          string  `json:"symbol,omitempty"`
		CoingeckoCoinID *string `json:"coingecko_coin_id,omitempty"`
	} `json:"attributes"`
}

// TokenDataItem
type TokenDataItem struct {
	basicStruct
	Attributes    TokenAttributesItem `json:"attributes"`
	Relationships struct {
		TopPools map[string][]basicStruct `json:"top_pools"`
	} `json:"relationships"`
}

// TokenAttributesItem
type TokenAttributesItem struct {
	Address           string            `json:"address"`
	Name              string            `json:"name"`
	Symbol            string            `json:"symbol"`
	CoingeckoCoinID   string            `json:"coingecko_coin_id"`
	Decimals          int               `json:"decimals"`
	TotalSupply       string            `json:"total_supply"`
	PriceUSD          string            `json:"price_usd"`
	FDVUSD            string            `json:"fdv_usd"`
	TotalReserveInUSD string            `json:"total_reserve_in_usd"`
	VolumeUSD         map[string]string `json:"volume_usd"`
	MarketCapUSD      *string           `json:"market_cap_usd"`
}

// TokenInfoDataItem
type TokenInfoDataItem struct {
	basicStruct
	Attributes    TokenInfoAttributesItem `json:"attributes"`
	Relationships map[string]basicStruct  `json:"relationships"`
}

// TokenInfoAttributesItem
type TokenInfoAttributesItem struct {
	Address           string    `json:"address"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	CoingeckoCoinID   string    `json:"coingecko_coin_id"`
	ImageURL          string    `json:"image_url"`
	Websites          []string  `json:"websites"`
	Description       string    `json:"description"`
	GTScore           float64   `json:"gt_score"`
	DiscordURL        *string   `json:"discord_url"`
	TelegramHandle    *string   `json:"telegram_handle"`
	TwitterHandle     *string   `json:"twitter_handle"`
	MetadataUpdatedAt time.Time `json:"metadata_updated_at,omitempty"`
}

// TokensInfoIncludedItem
type TokensInfoIncludedItem struct {
	basicStruct
	Attributes struct {
		Name            string  `json:"name"`
		CoingeckoCoinID *string `json:"coingecko_coin_id,omitempty"`
	} `json:"attributes"`
}
