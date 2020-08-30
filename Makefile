protoc:
	protoc -I internal/port_domain/ \
	internal/port_domain/port_domain.proto \
	--gogoslick_out=plugins=grpc:internal/port_domain
