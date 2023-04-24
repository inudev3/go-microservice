package server

import (
	"context"
	"currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct{
	currency.UnimplementedCurrencyServer
	log hclog.Logger
}
func Newcurrency(log hclog.Logger)*Currency {
	return &Currency{log: log}
}
func (c *Currency)GetRate(ctx context.Context, req *currency.RateRequest)(*currency.RateResponse, error){
	c.log.Info("Handle Get Rate", "base", req.GetBase(), "destination", req.GetDestination())
	rate:=float32(req.Base/req.Destination)
	return &currency.RateResponse{Rate: rate}, nil
}
