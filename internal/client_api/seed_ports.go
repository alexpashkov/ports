package client_api

import (
	"encoding/json"
	"github.com/alexpashkov/ports/internal/port_domain"
	"github.com/pkg/errors"
	"io"
)

func SeedPorts(dec *json.Decoder, seed func(*port_domain.Port) error) error {
	retrievePort, err := buildPortRetriever(dec)
	if err != nil {
		return errors.Wrap(err, "failed to build retriever")
	}
	var p port_domain.Port
	for p, err = retrievePort(); err == nil; p, err = retrievePort() {
		err := seed(&p)
		if err != nil {
			return errors.Wrap(err, "failed to upsert a port")
		}
	}
	if err != io.EOF {
		return err
	}
	return nil
}
