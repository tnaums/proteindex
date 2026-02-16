package main

import (
	"fmt"
	"github.com/tnaums/proteindex/internal/dex"
)

func commandInspect(cfg *config, args ...string) error {
	protein := args[0]
	p, ok := cfg.proteindex[protein]
	if !ok {
		fmt.Println("you have not caught that protein")
		return nil
	}
	//fmt.Printf("type is: %T\n\n", p)
	dex.ParseBlastp(p)
	return nil
}
