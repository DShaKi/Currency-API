package server

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-hclog"
	v3 "github.com/imakiri/currencyapi-go/v3"

	protos "github.com/DShaKi/Currency-API/protos/currency"
)

type Currency struct {
	log hclog.Logger
	protos.UnimplementedCurrencyServer
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l, protos.UnimplementedCurrencyServer{}}
}

func (c *Currency) GetRate(ctx context.Context, rq *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "Base", rq.GetBase(), "Destination", rq.GetDestination())

	httpClient := http.DefaultClient
	currencyExchangeClient, err := v3.NewClient("frt0urxO7RVlVOBkDu4T1uIMwY5rUzKRhkAxmh5o", httpClient)
	if err != nil {
		c.log.Error("Connecting to currency exchange api failed:", err)
		return nil, err
	}
	latestCurrencyExchangeAPIRequest := v3.LatestRequest{From: rq.GetBase(), To: []string{rq.GetDestination()}}
	currencyExchangeResponse, err := currencyExchangeClient.Latest(&latestCurrencyExchangeAPIRequest)
	if err != nil {
		c.log.Error("Getting the latest currency exchange rate failed:", err)
		return nil, err
	}

	currencyExchangeResponseDetails := currencyExchangeResponse.Data[rq.GetDestination()]
	rate := currencyExchangeResponseDetails.Value

	return &protos.RateResponse{Rate: float32(rate)}, nil
}
