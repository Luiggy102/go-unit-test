package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Luiggy102/go-unit-test/models"
	"github.com/gorilla/mux"
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

// endpont mocking
// "/pokemon/{id}"
// controller.GetPokemon(w http.ResponseWriter, r *http.Request)
func TestGetPokemon(t *testing.T) {
	c := require.New(t)

	// Request is Get
	r, err := http.NewRequest("GET", "/pokemon/{id}", nil)
	c.NoError(err)
	// ResponseWriter as new recorder
	w := httptest.NewRecorder()

	// simulate vars
	vars := map[string]string{
		"id": "bulbasaur",
	}

	// set the vars to the URL
	r = mux.SetURLVars(r, vars)

	// test the function
	GetPokemon(w, r)

	// expected result
	var expected models.Pokemon
	expectedBodyResponse, err := os.ReadFile("./samples/api_response.json")
	c.NoError(err)
	err = json.Unmarshal(expectedBodyResponse, &expected)
	c.NoError(err)

	// actual result (in the ResponseWriter)
	// as json
	var result models.Pokemon
	err = json.Unmarshal(w.Body.Bytes(), &result)
	c.NoError(err)

	// compare
	c.Equal(http.StatusOK, w.Code)
	c.Equal(expected, result)
}
