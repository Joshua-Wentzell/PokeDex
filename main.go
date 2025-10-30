package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	m := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
		},
	}
	m["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp(m),
	}
	inputScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		inputScanner.Scan()
		userInput := inputScanner.Text()
		userInputClean := cleanInput(userInput)
		if len(userInputClean) > 0 {
			command := userInputClean[0]
			value, ok := m[command]
			if ok {
				err := value.callback()
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func cleanInput(text string) []string {
	return strings.Split(strings.ToLower(strings.TrimSpace(text)), " ")
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(m map[string]cliCommand) func() error {
	return func() error {
		usage := "Welcome to the Pokedex!\nUsage:\n\n"
		for key := range m {
			usage += m[key].name + ": " + m[key].description + "\n"
		}
		fmt.Println(strings.TrimSuffix(usage, "\n"))
		return nil
	}
}
