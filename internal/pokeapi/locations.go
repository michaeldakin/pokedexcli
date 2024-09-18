package pokeapi

import (
	"encoding/json"
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
