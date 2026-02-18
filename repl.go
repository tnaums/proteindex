package main

import (
	"bufio"
	"fmt"
	"github.com/tnaums/proteindex/internal/dex"
	"github.com/tnaums/proteindex/internal/proteinapi"
	"os"
	"strings"
)

type config struct {
	proteinapiClient proteinapi.Client
	proteindex       map[string]dex.Protein
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nWelcome to the Proteindex! \ntype 'help' for a list of commands.\n")	
	for {

		fmt.Print("Proteindex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

// Function cleanInput returns a slice of lowercase strings
// from an input string where each is a 'field' or 'word'
// from the input string. Used to clean and parse repl
// input.
func cleanInput(text string) []string {
	//	lowered := strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays this help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Proteindex",
			callback:    commandExit,
		},
		"blastp": {
			name:        "blastp <name> <sequence>",
			description: "Submit blastp query",
			callback:    commandSubmitGo,
		},
		"proteindex": {
			name:        "proteindex",
			description: "List all captured proteins",
			callback:    commandProteindex,
		},
		"inspect": {
			name:        "inspect <protein>",
			description: "Inspect blastp results for a captured protein",
			callback:    commandInspect,
		},
	}
}
