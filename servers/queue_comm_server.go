package servers

import (
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/loggers"
	"github.com/9299381/wego/servers/commons"
	"github.com/9299381/wego/servers/queues"
	"github.com/9299381/wego/servers/transports"
	"github.com/go-kit/kit/endpoint"
)

/**
redis queue
*/
type QueueCommServer struct {
	*queues.Server
}

func NewQueueCommServer() *QueueCommServer {
	config := (&configs.QueueConfig{}).Load()
	ss := &QueueCommServer{
		Server: queues.NewServer(&queues.Options{
			Prefix:      config.Prefix,
			Listen:      config.Listen,
			Interval:    config.Interval,
			UseNumber:   true,
			Concurrency: config.Concurrency,
		}),
	}
	ss.Logger = loggers.Log
	return ss
}

func (it *QueueCommServer) Route(name string, endpoint endpoint.Endpoint) {

	handler := &commons.CommHandler{
		Handler: transports.NewQueue(endpoint),
	}
	it.Register(name, handler)
}

func (it *QueueCommServer) Load() {

	//注册通用路由
}

func (it *QueueCommServer) Start() error {
	return it.Serve()

}
