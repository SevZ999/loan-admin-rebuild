package serve

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/frame/ant"
	_ "github.com/small-ek/antgo/frame/serve/gin"
	"loan-ext/boot/asynq"
	"loan-ext/boot/pools"
	"loan-ext/boot/router"
)

// LoadSrv Load Api service<加载API服务>
func LoadSrv() {
	gin.ForceConsoleColor()

	configPath := flag.String("config", "./config/config.toml", "Configuration file path")

	flag.Parse()

	eng := ant.New(*configPath).AddFunc(func() {
		asynq.Register()
		pools.Register()
	}).Serve(router.Load())
	defer asynq.CloseClient()
	defer eng.Close()
}
