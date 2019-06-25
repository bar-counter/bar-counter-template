package template

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	DefaultRunAddr   = ":38080"
	DefaultName      = "/template"
	DefaultApiPrefix = "/path"
)

type Cfg struct {
	RunAddr   string
	APIBase   string
	ApiPrefix string
}

var nowCfg *Cfg

func initDefaultCfg() *Cfg {
	cfg := Cfg{
		RunAddr:   DefaultRunAddr,
		APIBase:   DefaultName,
		ApiPrefix: DefaultApiPrefix,
	}
	return &cfg
}

func checkCfg(cfg *Cfg) {
	defaultCfg := initDefaultCfg()
	if cfg.RunAddr == "" {
		cfg.RunAddr = defaultCfg.RunAddr
	}
	if cfg.APIBase == "" {
		cfg.APIBase = defaultCfg.APIBase
	}
	if cfg.ApiPrefix == "" {
		cfg.ApiPrefix = defaultCfg.ApiPrefix
	}
	nowCfg = cfg
}

func Register(r *gin.Engine, cfg *Cfg) error {
	checkCfg(cfg)
	mGroup := r.Group(cfg.APIBase)
	{
		mGroup.GET(cfg.ApiPrefix, apiPrefix)
		fmt.Printf("register api at http://127.0.0.1%v\n", cfg.RunAddr)
	}
	return nil
}

func NowCfg() *Cfg {
	return nowCfg
}

func apiPrefix(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, message)
}
