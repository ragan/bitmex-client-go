package bitmex

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	Host   string `json:"host"`
	Path   string `json:"path"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
	cl     *http.Client
}

type valData interface {
	toValues() url.Values
}

func New(host, path, key, secret string) Client {
	return Client{host, path, key, secret, http.DefaultClient}
}

func NewTestNet(key, secret string) Client {
	return New("testnet.bitmex.com", "/api/v1", key, secret)
}

func (c *Client) do(method, res, q string, form valData,
	v interface{}) (errData *Error) {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = c.Host
	u.Path = c.Path + res
	u.RawQuery = q

	fmt.Printf("Request url: %s\n", u.String())
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return
	}
	if form != nil {
		req.Form = form.toValues()
	}
	var expires = strconv.FormatInt(time.Now().Unix()+5, 10)
	req.Header = map[string][]string{
		"api-expires": {expires},
		"api-key":     {c.Key},
		"api-signature": {sign(c.Secret, method, u.Path, u.RawQuery, expires,
			req.PostForm.Encode())},
	}
	resp, err := c.cl.Do(req)
	fmt.Printf("Response status code: %d.\n", resp.StatusCode)
	bytes, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(string(bytes))
		json.Unmarshal(bytes, &errData)
		return
	}
	json.Unmarshal(bytes, v)
	return
}

func sign(secret, method, path, query, nonce, postStr string) string {
	str := ""
	if "" == query {
		str = strings.ToUpper(method) + path + nonce + postStr
	} else {
		str = strings.ToUpper(method) + path + "?" + query + nonce + postStr
	}
	return calc(secret, str)
}

func calc(secret, payload string) string {
	sig := hmac.New(sha256.New, []byte(secret))
	sig.Write([]byte(payload))
	return hex.EncodeToString(sig.Sum(nil))
}
