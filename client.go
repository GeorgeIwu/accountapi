package accountapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type AccountApi struct {
	url    string
	client *http.Client
}

func NewAccountApi(baseUrl string) *AccountApi {
	return &AccountApi{baseUrl + "/v1/organisation/accounts", &http.Client{}}
}

func (c AccountApi) Add(data *AccountData) (*AccountData, error) {
	payload, err := json.Marshal(&Account{data})
	if err != nil {
		return nil, err
	}

	req, err := c.createRequest("POST", "", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AccountApi) Get(id string) (*AccountData, error) {
	req, err := c.createRequest("GET", "/"+id, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c AccountApi) Remove(account *AccountData) error {
	req, err := c.createRequest("DELETE", "/"+account.ID+"?version="+strconv.FormatInt(*account.Version, 10), nil)
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode == 404 {
		return errors.New("record " + account.ID + " does not exist")
	}

	return getStatus(res)
}

func (c AccountApi) createRequest(method string, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.url+path, body)
	if err != nil {
		return nil, err
	}

	date := time.Now().GoString()
	req.Header.Add("Date", date)
	req.Header.Add("Accept", "application/vnd.api+json")

	return req, nil
}

func (c AccountApi) makeRequest(req *http.Request) (*AccountData, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = getStatus(res)
	if err != nil {
		return nil, err
	}

	return getAccountData(res)
}

func getStatus(res *http.Response) error {
	if res.StatusCode < 400 {
		return nil
	}

	var errorBody ErrorData
	err := getBodyData(res, &errorBody)
	if err != nil {
		return err
	}

	return errors.New(errorBody.ErrorMessage)
}

func getAccountData(res *http.Response) (*AccountData, error) {
	var account Account
	err := getBodyData(res, &account)
	if err != nil {
		return nil, err
	}

	return account.Data, nil
}

func getBodyData(res *http.Response, v interface{}) error {
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}
