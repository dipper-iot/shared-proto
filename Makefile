.PHONY:

srv-model:
	protoc --experimental_allow_proto3_optional --go_out=. \
      --go-grpc_out=. \
      --go-grpc_opt=paths=import \
      --proto_path=. ./services/model/*.proto

srv-device:
	protoc --experimental_allow_proto3_optional --go_out=. \
      --go-grpc_out=. \
      --go-grpc_opt=paths=import \
      --proto_path=. ./services/device/*.proto

srv-user:
	protoc --experimental_allow_proto3_optional --go_out=. \
      --go-grpc_out=. \
      --go-grpc_opt=paths=import \
      --proto_path=. ./services/user/*.proto

srv-mail:
	protoc --experimental_allow_proto3_optional --go_out=. \
      --go-grpc_out=. \
      --go-grpc_opt=paths=import \
      --proto_path=. ./services/mail/*.proto

srv-otp:
	protoc --experimental_allow_proto3_optional --go_out=. \
      --go-grpc_out=. \
      --go-grpc_opt=paths=import \
      --proto_path=. ./services/otp/*.proto