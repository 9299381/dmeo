package codecs

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/9299381/wego/constants"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/tools/idwork"
)

func CommandDecodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(req.(string)), &mapResult)
	if err != nil {
		return nil, errors.New(constants.ErrJson)
	}
	return contracts.Request{
		Id:   idwork.ID(),
		Data: mapResult,
	}, nil
}

func CommandEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}
