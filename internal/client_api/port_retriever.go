package client_api

import (
	"encoding/json"
	"fmt"
	"github.com/alexpashkov/ports/internal/port_domain"
	"io"
)

type portRetriever func() (port_domain.Port, error)

func buildPortRetriever(dec *json.Decoder) (portRetriever, error) {
	if _, err := dec.Token(); err != nil {
		return nil, err
	}
	return func() (port_domain.Port, error) {
		var p port_domain.Port
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
