package main

import (
	"fmt"

	"github.com/bar-counter/template"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// TODO init test here
	cfg := &template.Cfg{}
	err := template.Register(r, cfg)
	if err != nil {
		fmt.Printf("test Register err %v\n", err)
		return
	}
	nowCfg := template.NowCfg()
	err = r.Run(nowCfg.RunAddr)
	if err != nil {
		fmt.Printf("run err %v\n", err)
		return
	}
}
