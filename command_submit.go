package main

import (
	"fmt"
	//	"github.com/tnaums/proteindex/internal/proteinapi"
)

func commandSubmit(cfg *config, args ...string) error {
	fmt.Println("Submit the query")
	fmt.Printf("Query is %s\n", args[0])
	rid := cfg.proteinapiClient.SubmitBlast(args[0])
	fmt.Printf("commandSubmit_RID is: %s", rid)
	fmt.Println()
	return nil
}
