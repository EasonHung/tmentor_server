package internal_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type InternalApiHandler struct {
	baseUrl string
}

func NewApiHandler(baseUrl string) InternalApiHandler {
	return InternalApiHandler{
		baseUrl: baseUrl,
	}
}

func (this *InternalApiHandler) Get(path string) (error, []byte) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", this.baseUrl+path, nil)
	if err != nil {
		return errors.WithStack(err), nil
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.WithStack(err), nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("[internal api] status error, base url: %s, status code: %d", this.baseUrl, resp.StatusCode)), nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err), nil
	}

	return nil, respBody
}

func (this *InternalApiHandler) Post(path string, reqBody any) ([]byte, error) {
	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(this.baseUrl+path, "application/json", bytes.NewBuffer(reqBodyJson))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("[internal api] status error, base url: %s, status code: %d", this.baseUrl, resp.StatusCode))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return respBody, nil
}
