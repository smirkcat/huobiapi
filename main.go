package huobiapi

import (
	"github.com/bitly/go-simplejson"
	"github.com/smirkcat/huobiapi/client"
	"github.com/smirkcat/huobiapi/market"
)

type JSON = simplejson.Json

type ParamsData = client.ParamData
type Market = market.Market
type Listener = market.Listener
type Client = client.Client

// SetWebsocket .
func SetWebsocket(url string) {
	market.Endpoint = url
}

// NewMarket 创建WebSocket版Market客户端
func NewMarket() (*market.Market, error) {
	return market.NewMarket()
}

// NewClient 创建RESTFul客户端
func NewClient(accessKeyId, accessKeySecret string) (*client.Client, error) {
	return client.NewClient(client.Endpoint, accessKeyId, accessKeySecret)
}

// NewClientWithURL .
func NewClientWithURL(url, accessKeyId, accessKeySecret string) (*client.Client, error) {
	return client.NewClient(url, accessKeyId, accessKeySecret)
}
