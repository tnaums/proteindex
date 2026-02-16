package main

import "fmt"

func commandProteindex(cfg *config, args ...string) error {
	fmt.Println("Your Proteindex:")
	for k, _ := range cfg.proteindex {
		fmt.Printf(" - %s\n", k)
	}
	
	return nil
}
