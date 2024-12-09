generate-project-api:
	protoc --proto_path api/project_v1 --go_out=pkg/api/project_v1 --go_opt=paths=source_relative --go-grpc_out=pkg/api/project_v1 --go-grpc_opt=paths=source_relative api/project_v1/project.proto

generate-access-api:
	protoc --proto_path api/access_v1 --go_out=pkg/api/access_v1 --go_opt=paths=source_relative --go-grpc_out=pkg/api/access_v1 --go-grpc_opt=paths=source_relative api/access_v1/access.proto