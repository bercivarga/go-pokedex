package commands

import (
	"encoding/json"
	"fmt"
	"github.com/bercivarga/go-pokedex/state"
	"io"
	"net/http"
	"os"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(state *state.AppState, input string) error
}

var CommandMap = map[string]cliCommand{
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback: func(_ *state.AppState, _ string) error {
			fmt.Println("Exiting Pokedex...")
			os.Exit(0)
			return nil
		},
	},
	"joke": {
		Name:        "joke",
		Description: "Tell a joke",
		Callback: func(_ *state.AppState, _ string) error {
			fmt.Println("Why did the Squirtle cross the road? To get to the shell station!")
			return nil
		},
	},
	"map": {
		Name:        "map",
		Description: "Explore the location areas. Use multiple times to expand the list.",
		Callback: func(state *state.AppState, _ string) error {
			areas, err := getLocationAreas(state.GetAreaPage())
			if err != nil {
				return err
			}

			state.LocationAreas = areas
			state.AreaPage++

			fmt.Printf("Here is page %d of the location areas:\n", state.AreaPage)
			// print the location areas line by line
			for _, area := range state.LocationAreas {
				fmt.Println(area)
			}

			return nil
		},
	},
	"area": {
		Name:        "area",
		Description: "Get a location area",
		Callback: func(state *state.AppState, area string) error {
			pokemon, err := getPokemonInLocationArea(area)
			if err != nil {
				return err
			}

			fmt.Printf("Here are the Pokemon in the %s area:\n", area)
			// print the Pokemon line by line
			for _, p := range pokemon {
				fmt.Println(p)
			}

			return nil
		},
	},
	"pokeinfo": {
		Name:        "pokeinfo",
		Description: "Get information about a Pokemon",
		Callback: func(_ *state.AppState, name string) error {
			info, err := getPokemonInfo(name)
			if err != nil {
				return err
			}

			fmt.Println(info)

			return nil
		},
	},
}

func getLocationAreas(page int) ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%d&limit=20", page*20), nil)
	if err != nil {
		return []string{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []string{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	var locationAreas []string

	decoder := json.NewDecoder(resp.Body)
	for {
		var result map[string]interface{}
		if err := decoder.Decode(&result); err != nil {
			break
		}

		for _, locationArea := range result["results"].([]interface{}) {
			locationAreas = append(locationAreas, locationArea.(map[string]interface{})["name"].(string))
		}
	}

	return locationAreas, nil
}

func getPokemonInLocationArea(locationArea string) ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locationArea), nil)
	if err != nil {
		return []string{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []string{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	var pokemon []string

	decoder := json.NewDecoder(resp.Body)
	for {
		var result map[string]interface{}
		if err := decoder.Decode(&result); err != nil {
			break
		}

		for _, pokemonInArea := range result["pokemon_encounters"].([]interface{}) {
			pokemon = append(pokemon, pokemonInArea.(map[string]interface{})["pokemon"].(map[string]interface{})["name"].(string))
		}
	}

	return pokemon, nil
}

// Returns the pokemon info in pretty printed JSON format
func getPokemonInfo(name string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name), nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	var result map[string]interface{}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return "", err
	}

	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return "", err
	}

	return string(prettyJSON), nil
}
