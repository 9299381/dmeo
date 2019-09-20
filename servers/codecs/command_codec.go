package codecs

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
)

func CommandDecodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(req.(string)), &mapResult)
	if err !=nil{
		return nil, errors.New("args参数json解析错误")
	}
	return contracts.Request{
		Id:   wego.ID(),
		Data: mapResult,
	},nil
}

func CommandEncodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp,nil
}