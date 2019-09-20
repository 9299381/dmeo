package filters

import (
	"context"
	"errors"
	"github.com/9299381/wego"
	"github.com/9299381/wego/constants"
	"github.com/9299381/wego/contracts"
	"github.com/go-kit/kit/endpoint"
)

type ResponseEndpoint struct {
	next endpoint.Endpoint
}

func (it *ResponseEndpoint) Next(next endpoint.Endpoint) contracts.IFilter {
	it.next = next
	return it
}

func (it *ResponseEndpoint) Make() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		if wego.App.Status ==false{
			return contracts.ResponseFaile(errors.New(constants.ErrStop)),nil
		}
		resp,err := it.next(ctx, request)
		if err!= nil{
			return contracts.ResponseFaile(err),nil
		}
		return resp ,nil
	}
}