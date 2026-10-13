package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(name string) (DeepLocationAreaResp, error) {
	url := baseURL + "/location-area/" + name
	val, exists := c.cache.Get(url)
	if exists {
		var deepAreaResp DeepLocationAreaResp
		err := json.Unmarshal(val, &deepAreaResp)
		if err != nil {
			return DeepLocationAreaResp{}, err
		}
		return deepAreaResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return DeepLocationAreaResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return DeepLocationAreaResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return DeepLocationAreaResp{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return DeepLocationAreaResp{}, err
	}

	var deepAreaResp DeepLocationAreaResp
	err = json.Unmarshal(data, &deepAreaResp)
	if err != nil {
		return DeepLocationAreaResp{}, err
	}

	c.cache.Add(url, data)

	return deepAreaResp, nil
}
