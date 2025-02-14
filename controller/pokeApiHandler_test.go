package controller

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/Luiggy102/go-unit-test/models"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

// test for getting pokemon from api
func TestGetPokemonFromApi(t *testing.T) {
	var c = require.New(t)
	var expected, result models.PokeApiPokemonResponse

	body, err := os.ReadFile("./samples/poke_api_read.json")
	c.NoError(err)
	err = json.Unmarshal(body, &expected)
	c.NoError(err)

	result, err = GetPokemonFromPokeApi("bulbasaur")
	c.NoError(err)

	c.Equal(expected, result)
}

// test for mocking the http client
func TestGetPokemonFromApiWithMocks(t *testing.T) {
	var c = require.New(t)

	// http mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// test httpmock
	id := "bulbasaur"
	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)
	body, err := os.ReadFile("./samples/poke_api_read.json")
	c.NoError(err)

	httpmock.RegisterResponder("GET", request, // method
		httpmock.NewStringResponder(200, string(body))) // expected

	var result models.PokeApiPokemonResponse
	result, err = GetPokemonFromPokeApi("bulbasaur") // result
	c.NoError(err)

	var expected models.PokeApiPokemonResponse
	err = json.Unmarshal([]byte(body), &expected)
	c.NoError(err)

	c.Equal(expected, result)
}
