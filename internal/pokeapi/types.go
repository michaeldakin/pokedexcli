package pokeapi

// https://pokeapi.co/docs/v2#location-areas
// https://mholt.github.io/json-to-go/

// Location with pagination
type Locations struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Entire LocationArea
type PokemonArea struct {
	ID                   int                    `json:"id,omitempty"`
	Name                 string                 `json:"name,omitempty"`
	GameIndex            int                    `json:"game_index,omitempty"`
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates,omitempty"`
	Location             Location               `json:"location,omitempty"`
	Names                []Names                `json:"names,omitempty"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters,omitempty"`
}
type EncounterMethod struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type Version struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type EncounterVersionDetails struct {
	Rate    int     `json:"rate,omitempty"`
	Version GameVersion `json:"version,omitempty"`
}
type EncounterMethodRates struct {
	EncounterMethod EncounterMethod  `json:"encounter_method,omitempty"`
	VersionDetails  []EncounterVersionDetails `json:"version_details,omitempty"`
}
type Location struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type Language struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type Names struct {
	Name     string   `json:"name,omitempty"`
	Language Language `json:"language,omitempty"`
}
type EncounterPokemon struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type Method struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type EncounterDetails struct {
	MinLevel        int    `json:"min_level,omitempty"`
	MaxLevel        int    `json:"max_level,omitempty"`
	ConditionValues []any  `json:"condition_values,omitempty"`
	Chance          int    `json:"chance,omitempty"`
	Method          Method `json:"method,omitempty"`
}
type VersionDetails struct {
	Version          GameVersion            `json:"version,omitempty"`
	MaxChance        int                `json:"max_chance,omitempty"`
	EncounterDetails []EncounterDetails `json:"encounter_details,omitempty"`
}
type PokemonEncounters struct {
	Pokemon        EncounterPokemon          `json:"pokemon,omitempty"`
	VersionDetails []EncounterVersionDetails `json:"version_details,omitempty"`
}
