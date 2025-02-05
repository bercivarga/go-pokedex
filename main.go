package main

import (
	"fmt"
	"github.com/bercivarga/go-pokedex/commands"
	"github.com/bercivarga/go-pokedex/repl"
	"github.com/bercivarga/go-pokedex/state"
)

func main() {
	// Create a new appState
	appState := state.NewAppState()

	// Initialization message
	welcomeMessage := "Welcome to the Pokedex CLI!\nHere are the available commands:\n"

	for commandName, command := range commands.CommandMap {
		welcomeMessage += fmt.Sprintf("'%s': %s\n", commandName, command.Description)
	}

	fmt.Println(welcomeMessage)

	scanner := repl.CreateTerminalScanner()
	// Start an infinite loop
	for {
		// Print the prompt
		fmt.Print("Pokedex > ")
		// Wait for user input
		scanner.Scan()
		// Get the input
		input := scanner.Text()

		// Clean the input
		cleanedInput := repl.CleanInput(input)

		firstInput := cleanedInput[0]
		var secondInput string

		if len(cleanedInput) > 1 {
			secondInput = cleanedInput[1]
		}

		// Check if one of the commands is in the input
		if command, ok := commands.CommandMap[firstInput]; ok {
			// Run the command
			err := command.Callback(appState, secondInput)
			if err != nil {
				fmt.Printf("Error running command %s: %v\n", command.Name, err)
			}

			// Add a divider
			fmt.Println("--------------------------------------------------")

			// Skip the rest of the loop
			continue
		}
	}
}
