package builder

import (
	"decert/internal/pkg/config"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewServer() (*gin.Engine, error) {
	gin.SetMode(config.Instance().Http.Mode)
	server := gin.Default()
	setCORS(server)
	return server, nil
}

func setCORS(engine *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowMethods(http.MethodOptions)
	corsConfig.AllowAllOrigins = true
	engine.Use(cors.New(corsConfig))
}
