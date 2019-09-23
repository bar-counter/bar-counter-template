package barcountertemplate

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(r *gin.Engine, cfg *Cfg) error {
	mGroup := r.Group(cfg.APIBase)
	{
		mGroup.GET(cfg.ApiPrefix, apiPrefix)
		fmt.Printf("register api at http://127.0.0.1%v\n", cfg.RunAddr)
	}
	return nil
}

func apiPrefix(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, message)
}
