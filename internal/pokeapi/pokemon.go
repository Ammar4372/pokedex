package pokeapi

import (
	"encoding/json"
	"io"
)

func (api *Client) GetPokemon(name string) (Pokemon, error) {
	pageUrl := baseUrl + "pokemon/" + name

	if val, ok := api.cache.Get(pageUrl); ok {
		result := Pokemon{}
		err := json.Unmarshal(val, &result)
		if err != nil {
			return result, err
		}
		return result, nil
	}

	resp, err := api.httpClient.Get(pageUrl)
	data := Pokemon{}
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
