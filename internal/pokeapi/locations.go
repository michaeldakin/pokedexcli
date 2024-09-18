package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (Locations, error) {
	// Pagination by default
	// https://pokeapi.co/docs/v2#resource-listspagination-section
	//
	// Returns
	// count     int
	// next      string
	// previous: *string (nullable)
	// results   []results
	// - name
	// - url

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locations := Locations{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return Locations{}, err
		}

		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}

	locations := Locations{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return Locations{}, err
	}

	c.cache.Add(url, data)
	return locations, nil
}

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

