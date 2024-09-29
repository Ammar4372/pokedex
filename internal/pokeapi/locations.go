package pokeapi

import (
	"encoding/json"
	"io"
)

func (api *Client) GetLocations(url *string) (Locations, error) {
	pageUrl := baseUrl + "location-area/"
	if url == nil {
		url = &pageUrl
	}
	if val, ok := api.cache.Get(*url); ok {
		result := Locations{}
		err := json.Unmarshal(val, &result)
		if err != nil {
			return result, err
		}
		return result, nil
	}

	resp, err := api.httpClient.Get(*url)
	data := Locations{}
	if err != nil {
		return data, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	api.cache.Add(*url, body)
	return data, nil
}

func (api *Client) GetLocation(name string) (PokemonEncounters, error) {
	pageUrl := baseUrl + "location-area/" + name

	if val, ok := api.cache.Get(pageUrl); ok {
		result := PokemonEncounters{}
		err := json.Unmarshal(val, &result)
		if err != nil {
			return result, err
		}
		return result, nil
	}

	resp, err := api.httpClient.Get(pageUrl)
	data := PokemonEncounters{}
	if err != nil {
		return data, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}
	api.cache.Add(pageUrl, body)
	return data, nil

}
