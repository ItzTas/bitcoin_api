package client

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetCoinList() (*[]Coin, error) {
	url := BaseCoinDataURL + "list"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-cg-demo-api-key", c.geckoKey)

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	cd := []Coin{}
	if err = json.Unmarshal(dat, &cd); err != nil {
		return nil, err
	}

	return &cd, nil

}
