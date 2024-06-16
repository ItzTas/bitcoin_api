package client

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetCoinData(coinID string) (CoinData, error) {
	url := BaseCoinDataURL + coinID
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CoinData{}, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-cg-demo-api-key", c.geckoKey)

	res, err := c.Client.Do(req)
	if err != nil {
		return CoinData{}, err
	}
	defer res.Body.Close()
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return CoinData{}, err
	}

	cd := CoinData{}
	if err = json.Unmarshal(dat, &cd); err != nil {
		return CoinData{}, err
	}

	return cd, nil
}
