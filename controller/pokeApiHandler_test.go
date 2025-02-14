package controller

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/Luiggy102/go-unit-test/models"
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
