package wyre

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type client struct {
	http                 *http.Client
	key, secret, baseURL string
}

func NewClient(key, secret string, sandbox bool) *client {

	// get the base URL, based on sandbox flag
	baseURL := "https://api.sendwyre.com"
	if sandbox {
		baseURL = "https://api.testwyre.com"
	}
	baseURL += "/v2"

	return &client{
		http:    &http.Client{},
		key:     key,
		secret:  secret,
		baseURL: baseURL,
	}
}

func (c *client) makeURL(path string, params url.Values) string {
	return fmt.Sprintf("%s%s?%s", c.baseURL, normalizePath(path), params.Encode())
}

func (c *client) calculateRequestSignature(uri string, body []byte) (string, error) {
	data := fmt.Sprintf("%s%s", uri, body)

	h := hmac.New(sha256.New, []byte(c.secret))

	_, err := h.Write([]byte(data))
	if err != nil {
		return "", err
	}

	sigBytes := h.Sum(nil)
	return hex.EncodeToString(sigBytes), nil

}

func (c *client) doRequest(path, method string, params url.Values, body []byte, result interface{}) error {
	if params == nil {
		params = make(url.Values)
	}
	params.Add("timestamp", getTimestampString())
	uri := c.makeURL(path, params)

	req, err := http.NewRequest(method, uri, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	sig, err := c.calculateRequestSignature(uri, body)
	if err != nil {
		return err
	}

	req.Header.Add("X-Api-Key", c.key)
	req.Header.Add("X-Api-Signature", sig)

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if resp.StatusCode != 200 {
		apiErr := &APIError{}
		err = dec.Decode(apiErr)
		if err != nil {
			return err
		}
		return apiErr
	}

	err = dec.Decode(result)

	if err != nil {
		return err
	}

	return nil

}
