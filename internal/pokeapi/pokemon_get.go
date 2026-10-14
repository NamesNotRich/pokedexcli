package pokeapi

import(
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	val, exists := c.cache.Get(url)
	if exists {
		var pokeMon Pokemon
		err := json.Unmarshal(val, &pokeMon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokeMon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokeMon Pokemon
	err = json.Unmarshal(data, &pokeMon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return pokeMon, nil
}
