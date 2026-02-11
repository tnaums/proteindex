package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Proteindex... Goodbye!")
	os.Exit(0)
	return nil
}
