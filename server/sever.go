package server

import (
	"context"

	"github.com/hashicorp/go-hclog"

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
	c.log.Info("Handle GetRate", "base", rq.GetBase(), "destination", rq.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}
