package port_domain

import (
	"errors"
	"sync"
)

func NewService() *Service {
	return &Service{
		ports: make(map[string]*Port),
	}
}

type Service struct {
	ports map[string]*Port
	mu    sync.RWMutex
}

func (s *Service) UpsertPort(port *Port) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := portID(port)
	if id == "" {
		return errors.New("couldn't get port ID")
	}
	s.ports[id] = port
	return nil
}

func portID(p *Port) string {
	if p == nil || len(p.Unlocs) < 1 {
		return ""
	}
	return p.Unlocs[0]
}

func (s *Service) GetPortByID(id string) *Port {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.ports[id]
}
