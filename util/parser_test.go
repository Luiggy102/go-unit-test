package util

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/Luiggy102/go-unit-test/models"
	"github.com/stretchr/testify/require"
)

func TestParsePokemon(t *testing.T) {
	// require for assertions
	c := require.New(t)

	// get the poke-api response
	var pokeapiResponse models.PokeApiPokemonResponse
	// parse json from samples
	body, err := os.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)
	err = json.Unmarshal(body, &pokeapiResponse)
	c.NoError(err)

	// -- check the parse funcion -- //
	result, err := ParsePokemon(pokeapiResponse)
	c.NoError(err)

	// get the expected result
	var expected models.Pokemon
	body, err = os.ReadFile("./samples/api_response.json")
	c.NoError(err)
	err = json.Unmarshal(body, &expected)
	c.NoError(err)

	// compare
	c.Equal(expected, result)
}
