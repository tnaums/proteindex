package main

import (
	"fmt"
)

func commandSubmit(cfg *config, args ...string) error {
	protein := args[0]
	query := args[1]
	fmt.Printf("Submiting %s\n", protein)
	fmt.Printf("Query is %s\n", query)

	// if _, ok := cfg.proteinapiClient.Cache.Get(args[1]); ok {
	// 	fmt.Println("blastp results already in cache")
	// 	return nil
	// }
	
	rid, err := cfg.proteinapiClient.SubmitBlast(protein, query)
	if err != nil {
		fmt.Println(err)
	}
	if rid == "foundit" {
 	fmt.Println("blastp results already in cache")
		return nil
	}
	
	err = cfg.proteinapiClient.CheckBlast(protein, query, rid)
	return nil
}
