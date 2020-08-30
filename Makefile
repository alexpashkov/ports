protoc:
	protoc -I internal/port_domain/ \
	internal/port_domain/port_domain.proto \
	--go_out=plugins=grpc:internal/port_domain
