package wyre

import (
	"fmt"
)

type Rates map[string]float64

func (r Rates) GetPair(pair string) (float64, error) {
	if rate, ok := r[pair]; ok {
		return rate, nil
	}
	return 0, fmt.Errorf("%s is not a vlid exchange rate pair", pair)
}

func (c *Client) LiveExchangeRates() (Rates, error) {
	result := make(Rates)

	err := c.doRequest("rates", "GET", nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil

}
