package main

import (
	"time"

	"github.com/tnaums/proteindex/internal/proteinapi"
	"github.com/tnaums/proteindex/internal/dex"	
)

func main() {
	proteinClient := proteinapi.NewClient(5 * time.Second, time.Minute*5)
	proteinDex := dex.NewDex()
	cfg := &config{
		proteinapiClient: proteinClient,
		proteindex: proteinDex,
	}

	startRepl(cfg)
}
