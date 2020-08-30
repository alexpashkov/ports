package main

import (
	"context"
	"encoding/json"
	"github.com/alexpashkov/ports/internal/client_api"
	"github.com/alexpashkov/ports/internal/port_domain"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func main() {
	config, err := readConfigFromEnv()
	if err != nil {
		log.Fatalf("error reading config: %s", err)
	}
	conn, err := grpc.Dial(config.PortDomainAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %s", err)
	}
	portDomainClient := port_domain.NewPortDomainClient(conn)
	go mustSeedPorts(config.PortsFile, func(port *port_domain.Port) error {
		ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
		defer cancel()
		_, err := portDomainClient.UpsertPort(ctx, port)
		return err
	})
	s := &client_api.Service{
		Config:           config,
		PortDomainClient: portDomainClient,
	}
	err = s.Init()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to init the service"))
	}
}

func mustSeedPorts(portsFile string, seed func(*port_domain.Port) error) {
	f, err := os.Open(portsFile)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to open ports file"))
	}
	defer f.Close()
	err = client_api.SeedPorts(json.NewDecoder(f), seed)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to seed ports"))
	}
}

func readConfigFromEnv() (client_api.Config, error) {
	var (
		c = client_api.Config{
			Timeout: time.Second * 10,
		}
		ok bool
	)
	c.PortDomainAddress, ok = os.LookupEnv("PORT_DOMAIN_ENDPOINT")
	if !ok {
		return c, errors.New("PORT_DOMAIN_ENDPOINT is not set")
	}
	c.PortsFile, ok = os.LookupEnv("PORTS_FILE")
	if !ok {
		return c, errors.New("PORTS_FILE is not set")
	}
	rawTimeout, ok := os.LookupEnv("TIMEOUT")
	if ok {
		var err error
		c.Timeout, err = time.ParseDuration(rawTimeout)
		if err != nil {
			return c, errors.Wrapf(err, "failed to parse timeout: %q", rawTimeout)
		}
	}
	return c, nil
}
