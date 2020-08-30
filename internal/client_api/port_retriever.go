package client_api

import (
	"encoding/json"
	"fmt"
	"io"
)

type portRetriever func() (Port, error)

func buildPortRetriever(dec *json.Decoder) (portRetriever, error) {
	if _, err := dec.Token(); err != nil {
		return nil, err
	}
	return func() (Port, error) {
		var p Port
		if dec.More() {
			_, err := dec.Token()
			if err != nil {
				return p, err
			}
			if err := dec.Decode(&p); err != nil {
				return p, fmt.Errorf("failed to decode: %v", err)
			}
			return p, nil
		}
		if _, err := dec.Token(); err != nil {
			return p, err
		}
		return p, io.EOF
	}, nil
}
