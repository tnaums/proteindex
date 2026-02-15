package main

import (
	"fmt"
)

func commandSubmit(cfg *config, args ...string) error {
	fmt.Println("Submit the query")
	fmt.Printf("Query is %s\n", args[0])
	rid, err := cfg.proteinapiClient.SubmitBlast(args[0])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("commandSubmit_RID is: %s", rid)
	fmt.Println()
	return nil
}
