package server

import (
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
