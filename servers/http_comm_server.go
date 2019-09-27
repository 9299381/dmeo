package servers

import (
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/filters"
	"github.com/9299381/wego/loggers"
	"github.com/9299381/wego/servers/transports"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
)

type HttpCommServer struct {
	*mux.Router
	Logger contracts.ILogger
}

func NewHttpCommServer() *HttpCommServer {
	ss := &HttpCommServer{
		Router: mux.NewRouter(),
	}
	ss.Logger = loggers.Log
	return ss
}

func (it *HttpCommServer) Route(method string, path string, endpoint endpoint.Endpoint) {
	it.Methods(method).
		Path(path).
		Handler(transports.NewHTTP(endpoint))
}

func (it *HttpCommServer) Post(path string, endpoint endpoint.Endpoint) {
	it.Methods("POST").
		Path(path).
		Handler(transports.NewHTTP(endpoint))
}

//
func (it *HttpCommServer) Get(path string, endpoint endpoint.Endpoint) {
	it.Methods("GET").
		Path(path).
		Handler(transports.NewHTTP(endpoint))
}

func (it *HttpCommServer) Load() {

	//注册通用路由
	it.Route("GET", "/health", (&filters.HealthEndpoint{}).Make())

}

func (it *HttpCommServer) Start() error {
	config := (&configs.HttpConfig{}).Load()
	address := config.HttpHost + ":" + config.HttpPort
	it.Logger.Info("Http Server Start ", address)
	handler := it.Router
	return http.ListenAndServe(address, handler)
}
