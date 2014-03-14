package client

import (
	"bytes"
	"encoding/json"
)

func (client Client) Get(collection string, key string, value interface{}) error {
	resp, err := client.doRequest("GET", collection+"/"+key, nil)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return newError(resp)
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(value)

	return err
}

func (client Client) Put(collection string, key string, value interface{}) error {
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.Encode(value)

	resp, err := client.doRequest("PUT", collection+"/"+key, buf)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		err = newError(resp)
	}

	return err
}
