package wyre_test

import (
	"fmt"
	"github.com/dangersalad/gowyre"
	"testing"
	"time"
)

const API_KEY = "AK-8ZN7C9YZ-FYYT69T2-Z69RYJGH-BMALAVDQ"
const API_SECRET = "SK-NBYH7F6T-678WDA7L-NC8LT9QR-DCR73434"

func TestNewClient(t *testing.T) {
	c := wyre.NewClient(API_KEY, API_SECRET, true)
	if c == nil {
		t.Fatal("Client is nil")
	}
}

func TestRejectInvalidCredentials(t *testing.T) {
	c := wyre.NewClient(API_KEY, API_SECRET+"foobar", true)

	_, err := c.LiveExchangeRates()
	if err == nil {
		t.Fatal("API should reject bogus credentials!")
	}
	t.Log("Got error (expected)", err)
}

func TestLiveExchangeRates(t *testing.T) {
	c := wyre.NewClient(API_KEY, API_SECRET, true)

	rates, err := c.LiveExchangeRates()
	if err != nil {
		t.Fatal(err)
	}
	_, err = rates.GetPair("USDBTC")
	if err != nil {
		t.Fatal(err)
	}

	_, err = rates.GetPair("NOTREAL")
	if err == nil {
		t.Fatal("Rate map should error trying to find non existant key")
	}

}

func TestAccountInfo(t *testing.T) {
	c := wyre.NewClient(API_KEY, API_SECRET, true)

	account, err := c.AccountInfo()
	if err != nil {
		t.Fatal(err)
	}
	_, err = account.GetAddress("BTC")
	if err != nil {
		t.Fatal(err)
	}

	_, err = account.GetAddress("NOTREAL")
	if err == nil {
		t.Fatal("Account Info should error trying to find non existant address")
	}

}

func TestQuoteTransfer(t *testing.T) {
	c := wyre.NewClient(API_KEY, API_SECRET, true)

	rate, err := c.QuoteTransfer(&wyre.TransferQuoteRequest{
		SourceAmount:   100,
		SourceCurrency: "USD",
		DestCurrency:   "BTC",
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log("rate", rate)
}

func TestCreateTransfer(t *testing.T) {
	c := wyre.NewClient(API_KEY, API_SECRET, true)

	transfer, err := c.CreateTransfer(&wyre.TransferCreateRequest{
		SourceAmount:   100,
		SourceCurrency: "USD",
		DestCurrency:   "BTC",
		Dest:           "13zrYouP88q8UEcxG9nU47JwifaQjEJbhn",
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("transfer %#v", transfer)
}

func TestConfirmTransfer(t *testing.T) {
	c := wyre.NewClient(API_KEY, API_SECRET, true)

	transfer, err := c.CreateTransfer(&wyre.TransferCreateRequest{
		SourceAmount:   100,
		SourceCurrency: "USD",
		DestCurrency:   "BTC",
		Dest:           "13zrYouP88q8UEcxG9nU47JwifaQjEJbhn",
	})

	if err != nil {
		t.Fatal(err)
	}

	conf, err := c.ConfirmTransfer(transfer.ID)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("confirmed transfer %#v", conf)
}

func TestTransferStatus(t *testing.T) {
	c := wyre.NewClient(API_KEY, API_SECRET, true)

	transfer, err := c.CreateTransfer(&wyre.TransferCreateRequest{
		SourceAmount:   100,
		SourceCurrency: "USD",
		DestCurrency:   "BTC",
		Dest:           "13zrYouP88q8UEcxG9nU47JwifaQjEJbhn",
	})

	if err != nil {
		t.Fatal(err)
	}

	_, err = c.ConfirmTransfer(transfer.ID)
	if err != nil {
		t.Fatal(err)
	}

	transfers, err := c.TransferStatus(transfer.ID)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("transfer status %#v", transfers)
}

func TestTransferLookup(t *testing.T) {
	c := wyre.NewClient(API_KEY, API_SECRET, true)

	customId := fmt.Sprintf("foobar-%d", time.Now().Unix())

	transfer, err := c.CreateTransfer(&wyre.TransferCreateRequest{
		SourceAmount:   100,
		SourceCurrency: "USD",
		DestCurrency:   "BTC",
		Dest:           "13zrYouP88q8UEcxG9nU47JwifaQjEJbhn",
		CustomID:       customId,
	})

	if err != nil {
		t.Fatal(err)
	}

	_, err = c.ConfirmTransfer(transfer.ID)
	if err != nil {
		t.Fatal(err)
	}

	transfers, err := c.TransferLookup(customId)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("transfer lookup %#v", transfers)
}
