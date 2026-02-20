package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func commandFileSubmitGo(cfg *config, args ...string) error {
	seqFile, err := os.ReadDir("sequences") // []fs.DirEntry
	if err != nil {
		return err
	}

	for idx, seq := range seqFile {
		fmt.Printf(" %d. %s\n", idx+1, seq.Name())
	}

	var count int
	fmt.Print("Choose a sequence number: ")
	fmt.Scanln(&count)

	filename := seqFile[count-1].Name()
	data, err := os.ReadFile("sequences/" + filename)
	if err != nil {
		return err
	}

	extension := filepath.Ext(filename)
	proteinName := filename[0 : len(filename)-len(extension)]

	content := string(data)
	go commandSubmit(cfg, proteinName, content)
	return nil

}
