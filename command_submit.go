package main

import (
	"fmt"

	"github.com/tnaums/proteindex/internal/dex"
)

func commandSubmit(cfg *config, args ...string) error {
	protein := args[0]
	query := args[1]
	fmt.Printf("\nSubmiting %s\n", protein)
	fmt.Printf("\nQuery is %s\n", query)
	fmt.Print("Proteindex > ")	
	
	rid, err := cfg.proteinapiClient.SubmitBlast(protein, query)
	if err != nil {
		return err
	}
	if rid == "foundit" {
 	fmt.Println("blastp results already in cache")
		return nil
	}
	
	err = cfg.proteinapiClient.CheckBlast(protein, query, rid)
	if err != nil {
		return err
	}
	proteinData, err := cfg.proteinapiClient.CatchProtein(protein)
	if err != nil {
		return err
	}

	fmt.Printf("\nThrowing a Pokeball at %s...\n", protein)
	cfg.proteindex[protein] = dex.Protein{
		Name:  protein,
		Blast: proteinData,
	}
	fmt.Printf("Caught %s!\n", protein)
	fmt.Printf("You have caught %d proteins.\n", len(cfg.proteindex))
	fmt.Print("Proteindex > ")
	
	return nil
}
