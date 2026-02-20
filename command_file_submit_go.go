package main

import (
	"fmt"
	"os"
)

func commandFileSubmitGo(cfg *config, args ...string) error {
	seqFile, err := os.ReadDir("sequences")
	if err != nil {
		return err
	}
	for idx, seq := range seqFile {
		fmt.Printf(" %d. %s\n", idx + 1, seq.Name())
	}

	
	data, err := os.ReadFile("sequences/" + seqFile[0].Name())
	if err != nil {
		return err
	}
	content := string(data)
	go commandSubmit(cfg, args[0], content)
	return nil

}
