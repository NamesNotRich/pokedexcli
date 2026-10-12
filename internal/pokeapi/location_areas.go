package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchLocationAreas(pageURL *string) (LocationAreaResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	var locationAreaResp LocationAreaResp
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}
	return locationAreaResp, nil
}
