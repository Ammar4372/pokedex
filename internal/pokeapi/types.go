package pokeapi

type Location struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}
type Locations struct {
	Locations []Location `json:"results"`
	Next      string     `json:"next"`
	Previous  string     `json:"previous"`
}
type PokemonEncounters struct {
	Pokemons []Pokemons `json:"pokemon_encounters"`
}
type Pokemons struct {
	Pokemon Pokemon `json:"pokemon"`
}
type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
