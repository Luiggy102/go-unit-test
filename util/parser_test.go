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

// for no pokemon type in json
func TestParsePokemonTypeNotFound(t *testing.T) {
	// require for assertions
	c := require.New(t)

	// get the poke-api response
	var pokeapiResponse models.PokeApiPokemonResponse
	// parse json from samples
	body, err := os.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)
	err = json.Unmarshal(body, &pokeapiResponse)
	c.NoError(err)

	// empty pokemon type (panic error)
	pokeapiResponse.PokemonType = []models.PokemonType{}

	// -- check the parse funcion -- //
	_, err = ParsePokemon(pokeapiResponse)
	c.NotNil(err) // is an error

	// the same error
	c.EqualError(ErrNotFoundPokemonType, err.Error())
}

func BenchmarkParser(b *testing.B) {
	c := require.New(b)

	var pokeapiResponse models.PokeApiPokemonResponse
	body, err := os.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)
	err = json.Unmarshal(body, &pokeapiResponse)
	c.NoError(err)

	for n := 0; n < b.N; n++ {
		_, err := ParsePokemon(pokeapiResponse)
		c.NoError(err)
	}
}
