package main

import (
	"fmt"

	template "github.com/bar-counter/bar-counter-template"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// TODO init test here
	cfg := template.NewCfg()
	err := template.Register(r, cfg)
	if err != nil {
		fmt.Printf("test Register err %v\n", err)
		return
	}
	nowCfg := template.NewCfg(
		template.WithRunAddr(template.DefaultRunAddr),
	)
	err = r.Run(nowCfg.RunAddr)
	if err != nil {
		fmt.Printf("run err %v\n", err)
		return
	}
}
