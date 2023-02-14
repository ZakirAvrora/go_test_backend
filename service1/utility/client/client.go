package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type client struct {
	url string
}

func New(url string) *client {
	return &client{url: url}
}

func (c *client) PostReq() ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, c.url, nil)
	if err != nil {
		return nil, fmt.Errorf("request for service 1 is not correct: %w", err)
	}

	httpClient := &http.Client{}

	res, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("error in sending req to service 1 : %w", err)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
