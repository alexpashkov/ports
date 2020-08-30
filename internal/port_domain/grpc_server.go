package port_domain

import "context"

type GRPCServer struct {
	*Service
}

func (s *GRPCServer) UpsertPort(_ context.Context, port *Port) (*UpsertPortResponse, error) {
	return &UpsertPortResponse{}, s.Service.UpsertPort(port)
}

func (s *GRPCServer) GetPort(_ context.Context, request *GetPortRequest) (*GetPortResponse, error) {
	return &GetPortResponse{Port: s.Service.GetPortByID(request.Id)}, nil
}
