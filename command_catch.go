package main

import (
	"github.com/tnaums/proteindex/internal/dex"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	protein := args[0]
	proteinData, err := cfg.proteinapiClient.CatchProtein(protein)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", protein)
	cfg.proteindex[protein] = dex.Protein{
		Name:  protein,
		Blast: proteinData,
	}
	fmt.Printf("You have caught %d proteins.\n", len(cfg.proteindex))
	return nil
}
