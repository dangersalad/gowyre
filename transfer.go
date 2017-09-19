package wyre

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type TransferQuoteRequest struct {
	SourceCurrency string  `json:"sourceCurrency"`
	DestCurrency   string  `json:"destCurrency"`
	SourceAmount   float64 `json:"sourceAmount"`
}

type TransferCreateRequest struct {
	// required
	SourceCurrency string `json:"sourceCurrency"`
	Dest           string `json:"dest"`
	DestCurrency   string `json:"destCurrency"`

	// one of the following two are required, but not both
	SourceAmount float64 `json:"sourceAmount,omitempty"`
	DestAmount   float64 `json:"destAmount,omitempty"`

	// optional
	Source             string `json:"source,omitempty"`
	Message            string `json:"message,omitempty"`
	CallbackURL        string `json:"callbackUrl,omitempty"`
	AutoConfirm        bool   `json:"autoConfirm,omitempty"`
	CustomID           string `json:"customId,omitempty"`
	AmountIncludesFees bool   `json:"amountIncludesFees,omitempty"`
	Preview            bool   `json:"preview,omitempty"`
	MuteMessages       bool   `json:"muteMessages,omitempty"`
}

type Transfer struct {
	ID                      string             `json:"id"`
	Status                  string             `json:"status"`
	FailureReason           string             `json:"failureReason"`
	Language                string             `json:"language"`
	CreatedAt               int64              `json:"createdAt"`
	CompletedAt             int64              `json:"completedAt"`
	DepositInitiatedAt      int64              `json:"depositInitiatedAt"`
	CancelledAt             int64              `json:"cancelledAt"`
	ExpiresAt               int64              `json:"expiresAt"`
	Owner                   string             `json:"owner"`
	Source                  string             `json:"source"`
	Dest                    string             `json:"dest"`
	SourceCurrency          string             `json:"sourceCurrency"`
	SourceAmount            float64            `json:"sourceAmount"`
	DestCurrency            string             `json:"destCurrency"`
	DestAmount              float64            `json:"destAmount"`
	ExchangeRate            float64            `json:"exchangeRate"`
	Desc                    string             `json:"desc"`
	Message                 string             `json:"message"`
	TotalFees               float64            `json:"totalFees"`
	Equivalencies           map[string]float64 `json:"equivalencies"`
	FeeEquivalencies        map[string]float64 `json:"feeEquivalencies"`
	Fees                    map[string]float64 `json:"fees"`
	AuthorizingIP           string             `json:"authorizingIp"`
	PaymentUrl              string             `json:"paymentUrl"`
	ExchangeOrderID         string             `json:"exchangeOrderId"`
	ChargeID                string             `json:"chargeId"`
	DepositID               string             `json:"depositId"`
	SourceTxID              string             `json:"sourceTxId"`
	DestTxID                string             `json:"destTxId"`
	CustomID                string             `json:"customId"`
	Buy                     bool               `json:"buy"`
	InstantBuy              bool               `json:"instantBuy"`
	Sell                    bool               `json:"sell"`
	Exchange                bool               `json:"exchange"`
	Send                    bool               `json:"send"`
	Deposit                 bool               `json:"deposit"`
	Withdrawal              bool               `json:"withdrawal"`
	Closed                  bool               `json:"closed"`
	ReversingSubStatus      string             `json:"reversingSubStatus"`
	ReversalReason          string             `json:"reversalReason"`
	RetrievalUrl            string             `json:"retrievalUrl"`
	QuotedMargin            float64            `json:"quotedMargin"`
	PendingSubStatus        string             `json:"pendingSubStatus"`
	DestName                string             `json:"destName"`
	SourceName              string             `json:"sourceName"`
	ExchangeOrder           string             `json:"exchangeOrder"`
	EstimatedArrival        int64              `json:"estimatedArrival"`
	BlockchainTx            string             `json:"blockchainTx"`
	Documents               []interface{}      `json:"documents"`
	ReversalRevenue         float64            `json:"reversalRevenue"`
	ReversalRevenueCurrency string             `json:"reversalRevenueCurrency"`
	DepositInfo             string             `json:"depositInfo"`
	ChargeInfo              string             `json:"chargeInfo"`
}

func (c *client) QuoteTransfer(quote *TransferQuoteRequest) (float64, error) {
	result := make(map[string]float64)

	body, err := json.Marshal(quote)
	if err != nil {
		return 0, err
	}

	err = c.doRequest("quote", "POST", nil, body, &result)
	if err != nil {
		return 0, err
	}
	return result["rate"], nil

}

func (c *client) CreateTransfer(t *TransferCreateRequest) (*Transfer, error) {
	result := &Transfer{}

	body, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	err = c.doRequest("transfers", "POST", nil, body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) ConfirmTransfer(id string) (*Transfer, error) {
	result := &Transfer{}

	err := c.doRequest(fmt.Sprintf("transfer/%s/confirm", id), "POST", nil, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) TransferStatus(id string) (*Transfer, error) {
	result := &Transfer{}

	err := c.doRequest(fmt.Sprintf("transfer/%s", id), "GET", nil, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) TransferLookup(id string) (*Transfer, error) {
	result := &Transfer{}

	params := make(url.Values)
	params.Add("customId", id)

	err := c.doRequest("transfer", "GET", params, nil, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
