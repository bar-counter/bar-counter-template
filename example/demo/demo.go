package main

import (
	"fmt"

	"github.com/bar-counter/bar-counter-template"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// TODO init test here
	cfg := barcountertemplate.NewCfg()
	err := barcountertemplate.Register(r, cfg)
	if err != nil {
		fmt.Printf("test Register err %v\n", err)
		return
	}
	nowCfg := barcountertemplate.NewCfg(
		barcountertemplate.WithRunAddr(barcountertemplate.DefaultRunAddr),
	)
	err = r.Run(nowCfg.RunAddr)
	if err != nil {
		fmt.Printf("run err %v\n", err)
		return
	}
}
