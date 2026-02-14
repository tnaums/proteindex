package main

import (
	"time"

	"github.com/tnaums/proteindex/internal/proteinapi"
)

func main() {
	proteinClient := proteinapi.NewClient(5 * time.Second, time.Minute*5)
	cfg := &config{
		proteinapiClient: proteinClient,
	}

	startRepl(cfg)
}
