package client

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetCoinList(limit *int) (*[]Coin, error) {
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

	coins := []Coin{}
	if err = json.Unmarshal(dat, &coins); err != nil {
		return nil, err
	}

	if limit == nil {
		return &coins, nil
	}

	if *limit > len(coins) {
		*limit = len(coins)
	}

	cdFilter := make([]Coin, *limit)
	for i := range *limit {
		cdFilter[i] = coins[i]
	}

	return &cdFilter, nil

}
