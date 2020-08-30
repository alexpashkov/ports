package client_api

import (
	"github.com/alexpashkov/ports/internal/port_domain"
	"time"
)

type Service struct {
	Config
	port_domain.PortDomainClient
}

type Config struct {
	PortDomainAddress, PortsFile string
	Timeout                      time.Duration
}

func (s *Service) Init() error {
	return nil
}
