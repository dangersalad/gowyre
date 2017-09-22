package wyre

import (
	"fmt"
)

// limited set of the results returned from the /account endpoint,
// could be expanded if necessary
type AccountInfo struct {
	ID                string             `json:"id"`
	DepositAddresses  map[string]string  `json:"depositAddresses"`
	TotalBalances     map[string]float64 `json:"totalBalances"`
	AvailableBalances map[string]float64 `json:"availableBalances"`
}

func (a *AccountInfo) GetAddress(currency string) (string, error) {
	if addr, ok := a.DepositAddresses[currency]; ok {
		return addr, nil
	}
	return "", fmt.Errorf("%s is not a valid deposit address currency", currency)
}

func (c *Client) AccountInfo() (*AccountInfo, error) {
	result := &AccountInfo{}

	err := c.doRequest("account", "GET", nil, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil

}
