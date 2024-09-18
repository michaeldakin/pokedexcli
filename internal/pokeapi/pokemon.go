package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationPokemon(areaName string) (PokemonArea, error) {
    url := fmt.Sprintf("%s/location-area/%s/", baseURL, areaName)
    // fmt.Printf("DEBUG URL: %q, areaName: %q\n", url, areaName)

    if val, ok := c.cache.Get(url); ok {
        encounter := PokemonArea{}
        err := json.Unmarshal(val, &encounter)
        if err != nil {
            return PokemonArea{}, err
        }

        return encounter, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return PokemonArea{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return PokemonArea{}, err
    }
    defer res.Body.Close()

    data, err := io.ReadAll(res.Body)
    if err != nil {
        return PokemonArea{}, err
    }

    encounters := PokemonArea{}
    err = json.Unmarshal(data, &encounters)
    if err != nil {
        return PokemonArea{}, err
    }

    c.cache.Add(url, data)
    return encounters, nil
}
