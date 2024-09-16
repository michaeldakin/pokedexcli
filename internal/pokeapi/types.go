package pokeapi

// https://pokeapi.co/docs/v2#location-areas
// https://mholt.github.io/json-to-go/

type Locations struct {
	Count    int       `json:"count"`
	Next     *string    `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
