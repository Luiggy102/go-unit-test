package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Luiggy102/go-unit-test/models"
	"github.com/Luiggy102/go-unit-test/util"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	ErrPokemonNotFound = errors.New("pokemon not found")
	ErrPokeApiFailure  = errors.New("unexpected response in Pokeapi")
)

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	// separate functions
	apiPokemon, err := GetPokemonFromPokeApi(id)

	// check the errors with errors package
	if errors.Is(err, ErrPokemonNotFound) {
		respondwithJSON(w, http.StatusNotFound, fmt.Sprintf("pokemon not found: %s", id))
	}

	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, fmt.Sprintf("error while calling pokeapi: %s", err.Error()))
	}

	parsedPokemon, err := util.ParsePokemon(apiPokemon)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, fmt.Sprintf("error found: %s", err.Error()))
	}

	respondwithJSON(w, http.StatusOK, parsedPokemon)
}

func GetPokemonFromPokeApi(id string) (models.PokeApiPokemonResponse, error) {
	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	response, err := http.Get(request)
	if err != nil {
		return models.PokeApiPokemonResponse{}, err
	}

	// check the status code from the response
	if response.StatusCode == http.StatusNotFound {
		return models.PokeApiPokemonResponse{}, ErrPokemonNotFound
	}

	if response.StatusCode != http.StatusOK {
		return models.PokeApiPokemonResponse{}, ErrPokeApiFailure
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.PokeApiPokemonResponse{}, err
	}

	var apiPokemon models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &apiPokemon)
	if err != nil {
		return models.PokeApiPokemonResponse{}, err
	}

	return apiPokemon, nil
}
