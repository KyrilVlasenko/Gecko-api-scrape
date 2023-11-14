package coingecko

import (
	"context"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestClient_GetCirculatingSupplyChartByCoinID(t *testing.T) {
	cases := []struct {
		name         string
		server       *httptest.Server
		id           string
		wantedIsErr  bool
		wantedResult *CoinCirculatingSupplyChartResponse
		wantedErrStr string
	}{
		{
			name: "success",
			server: mockHTTPServer(t, &CoinCirculatingSupplyChartResponse{CirculatingSupply: []struct {
				CoinsIDCirculatingSupplyChartItem
			}{
				{
					CoinsIDCirculatingSupplyChartItem{
						"1666224000000",
						"19184000.0",
					},
				},
			}}),
			id:          "ethereum",
			wantedIsErr: false,
			wantedResult: &CoinCirculatingSupplyChartResponse{CirculatingSupply: []struct {
				CoinsIDCirculatingSupplyChartItem
			}{
				{
					CoinsIDCirculatingSupplyChartItem{
						"1666224000000",
						"19184000.0",
					},
				},
			}},
			wantedErrStr: "",
		},
		{
			name:         "empty id param",
			id:           "",
			server:       mockHTTPServer(t, nil),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "id should not be empty",
		},
		{
			name:         "failed to call api",
			id:           "ethereum",
			server:       mockErrorHTTPServer(t),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: statusCode400ErrStr,
		},
		{
			name:         "failed to unmarshal json",
			id:           "ethereum",
			server:       mockHTTPServer(t, []byte(`{"name":what?}`)),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: incorrectJSONTypeErrStr,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := setup(t)
			client.apiURL = tt.server.URL
			result, err := client.GetCirculatingSupplyChartByCoinID(context.TODO(), tt.id, 1, "daily")
			if tt.wantedIsErr {
				if !strings.Contains(err.Error(), tt.wantedErrStr) {
					t.Fatalf("incorrect error, wanted error: %v, got error: %v", tt.wantedErrStr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("error should be nil, got: %v", err)
				}
			}
			if !reflect.DeepEqual(result, tt.wantedResult) {
				t.Fatalf("incorrect response, wanted result: %+v, got result: %+v", tt.wantedResult, result)
			}
		})
	}
}

func TestClient_GetCirculatingSupplyChartRangeByCoinID(t *testing.T) {
	cases := []struct {
		name         string
		server       *httptest.Server
		id           string
		wantedIsErr  bool
		wantedResult *CoinCirculatingSupplyChartResponse
		wantedErrStr string
	}{
		{
			name: "success",
			server: mockHTTPServer(t, &CoinCirculatingSupplyChartResponse{CirculatingSupply: []struct {
				CoinsIDCirculatingSupplyChartItem
			}{
				{
					CoinsIDCirculatingSupplyChartItem{
						"1666224000000",
						"19184000.0",
					},
				},
			}}),
			id:          "ethereum",
			wantedIsErr: false,
			wantedResult: &CoinCirculatingSupplyChartResponse{CirculatingSupply: []struct {
				CoinsIDCirculatingSupplyChartItem
			}{
				{
					CoinsIDCirculatingSupplyChartItem{
						"1666224000000",
						"19184000.0",
					},
				},
			}},
			wantedErrStr: "",
		},
		{
			name:         "empty id param",
			id:           "",
			server:       mockHTTPServer(t, nil),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "id should not be empty",
		},
		{
			name:         "failed to call api",
			id:           "ethereum",
			server:       mockErrorHTTPServer(t),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: statusCode400ErrStr,
		},
		{
			name:         "failed to unmarshal json",
			id:           "ethereum",
			server:       mockHTTPServer(t, []byte(`{"name":what?}`)),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: incorrectJSONTypeErrStr,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := setup(t)
			client.apiURL = tt.server.URL
			result, err := client.GetCirculatingSupplyChartRangeByCoinID(context.TODO(), tt.id, 1633046400, 1635724799)
			if tt.wantedIsErr {
				if !strings.Contains(err.Error(), tt.wantedErrStr) {
					t.Fatalf("incorrect error, wanted error: %v, got error: %v", tt.wantedErrStr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("error should be nil, got: %v", err)
				}
			}
			if !reflect.DeepEqual(result, tt.wantedResult) {
				t.Fatalf("incorrect response, wanted result: %+v, got result: %+v", tt.wantedResult, result)
			}
		})
	}
}

func TestClient_ListAllTokensByAssetPlatformID(t *testing.T) {
	// resp := &ListAllTokensResponse{
	// 	Name:      "CoinGecko",
	// 	LogoURI:   "https://www.coingecko.com/assets/thumbnail-007177f3eca19695592f0b8b0eabbdae282b54154e1be912285c9034ea6cbaf2.png",
	// 	Keywords:  []string{"defi"},
	// 	Timestamp: time.Unix(1666224000, 0),
	// 	Tokens: []TokensListAllItem{
	// 		{
	// 			ChainID:  137,
	// 			Address:  "0xb33eaad8d922b1083446dc23f610c2567fb5180f",
	// 			Name:     "Uniswap",
	// 			Symbol:   "UNI",
	// 			Decimals: 18,
	// 			LogoURI:  "https://assets.coingecko.com/coins/images/12504/thumb/uniswap-uni.png?1600306604",
	// 		},
	// 	},
	// 	Version: struct {
	// 		Major int `json:"major"`
	// 		Minor int `json:"minor"`
	// 		Patch int `json:"patch"`
	// 	}{141, 4, 0},
	// }
	cases := []struct {
		name         string
		server       *httptest.Server
		id           string
		wantedIsErr  bool
		wantedResult *ListAllTokensResponse
		wantedErrStr string
	}{
		// {
		// 	name:         "success",
		// 	server:       mockHTTPServer(t, resp),
		// 	id:           "polygon-pos",
		// 	wantedIsErr:  false,
		// 	wantedResult: resp,
		// 	wantedErrStr: "",
		// },
		{
			name:         "empty id param",
			id:           "",
			server:       mockHTTPServer(t, nil),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "id should not be empty",
		},
		{
			name:         "failed to call api",
			id:           "ethereum",
			server:       mockErrorHTTPServer(t),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: statusCode400ErrStr,
		},
		{
			name:         "failed to unmarshal json",
			id:           "ethereum",
			server:       mockHTTPServer(t, []byte(`{"name":what?}`)),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: incorrectJSONTypeErrStr,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := setup(t)
			client.apiURL = tt.server.URL
			result, err := client.ListAllTokensByAssetPlatformID(context.TODO(), tt.id)
			if tt.wantedIsErr {
				if !strings.Contains(err.Error(), tt.wantedErrStr) {
					t.Fatalf("incorrect error, wanted error: %v, got error: %v", tt.wantedErrStr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("error should be nil, got: %v", err)
				}
			}
			if !reflect.DeepEqual(result, tt.wantedResult) {
				t.Fatalf("incorrect response, wanted result: %+v, got result: %+v", tt.wantedResult, result)
			}
		})
	}
}
