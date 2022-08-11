package builder

import (
	"decert/internal/pkg/api"
	"fmt"

	"github.com/gin-gonic/gin"
)

type apiServer struct {
	server *gin.Engine
}

func NewApiServer() (IRunner, error) {
	server, _ := NewServer()
	api.AddRouterV1(server)
	return &apiServer{server: server}, nil
}

func (f *apiServer) Run(bindAddress string) error {
	fmt.Println(bindAddress)
	return f.server.Run(bindAddress)
}
