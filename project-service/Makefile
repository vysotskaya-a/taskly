generate-project-api:
	protoc --proto_path api/project_v1 --go_out=pkg/api/project_v1 --go_opt=paths=source_relative --go-grpc_out=pkg/api/project_v1 --go-grpc_opt=paths=source_relative api/project_v1/project.proto

generate-task-api:
	protoc --proto_path api/task_v1 --go_out=pkg/api/task_v1 --go_opt=paths=source_relative --go-grpc_out=pkg/api/task_v1 --go-grpc_opt=paths=source_relative api/task_v1/task.proto
