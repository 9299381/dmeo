package servers

import (
	"context"
	"github.com/9299381/wego"
	"github.com/9299381/wego/args"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/filters"
	"github.com/9299381/wego/loggers"
	"github.com/9299381/wego/servers/transports"
	"github.com/9299381/wego/servers/transports/codecs"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
	"os/exec"
)

type HttpCommServer struct {
	*mux.Router
	Logger contracts.ILogger
}

func NewHttpCommServer() *HttpCommServer {
	ss := &HttpCommServer{
		Router: mux.NewRouter(),
	}
	ss.Logger = loggers.GetLog()
	return ss
}

func (s *HttpCommServer) Route(method string, path string, endpoint endpoint.Endpoint) {
	s.Methods(method).
		Path(path).
		Handler(transports.NewHTTP(endpoint))
}

func (s *HttpCommServer) Post(path string, endpoint endpoint.Endpoint) {
	s.Methods("POST").
		Path(path).
		Handler(transports.NewHTTP(endpoint))
}

//
func (s *HttpCommServer) Get(path string, endpoint endpoint.Endpoint) {
	s.Methods("GET").
		Path(path).
		Handler(transports.NewHTTP(endpoint))
}

func (s *HttpCommServer) Load() {

	//注册通用路由
	s.Route("GET", "/health", (&filters.HealthEndpoint{}).Make())
	if args.Mode != "prod" {
		s.handleSwagger()
	}
}

func (s *HttpCommServer) Start() error {
	config := (&configs.HttpConfig{}).Load()
	address := config.HttpHost + ":" + config.HttpPort
	s.Logger.Info("Http Server Start ", address)
	handler := s.Router
	return http.ListenAndServe(address, handler)
}

func (s *HttpCommServer) handleSwagger() {
	//文件服务器 /swagger/ 前缀目录下index.html
	fs := http.FileServer(http.Dir("./swaggerui/"))
	s.Methods("GET").
		PathPrefix("/swagger/").
		Handler(http.StripPrefix("/swagger/", fs))
	//重新生成
	s.Methods("GET").
		Path("/swagger_generate").
		HandlerFunc(s.buildSwagger())
}

func (s *HttpCommServer) buildSwagger() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command(
			"swagger",
			"generate", "spec", "-o", "swaggerui/swagger.json")
		err := cmd.Run()
		response := contracts.MakeResponse("ok", err)
		_ = codecs.HttpEncodeResponse(context.Background(), w, response)
	}
}

func (s *HttpCommServer) Close() {
	v, ok := wego.App.Consul["http"]
	if ok {
		v.Deregister()
	}
}
