include .env


LOCAL_BIN:=$(CURDIR)/bin
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD)"
LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/protoc-gen-validate


get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-auth-user-api
	make generate-access-api
	make generate-auth-api



generate-auth-api:
	mkdir "pkg/auth_v1"
	protoc --proto_path api/auth_v1 \
        	--go_out=pkg/auth_v1 --go_opt=paths=source_relative \
        	--go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
        	api/auth_v1/auth_v1.proto

generate-access-api:
	mkdir "pkg/access_v1"
	protoc --proto_path api/access_v1 \
    	--go_out=pkg/access_v1 --go_opt=paths=source_relative \
    	--go-grpc_out=pkg/access_v1 --go-grpc_opt=paths=source_relative \
    	api/access_v1/access_v1.proto


generate-auth-user-api:
	mkdir "pkg/auth_user_v1"
	protoc --proto_path api/auth_user_v1 --proto_path vendor.protogen \
	--go_out=pkg/auth_user_v1 --go_opt=paths=source_relative \
	--go-grpc_out=pkg/auth_user_v1 --go-grpc_opt=paths=source_relative \
	--validate_out lang=go:pkg/auth_user_v1 --validate_opt=paths=source_relative \
	api/auth_user_v1/auth_user_v1.proto

create-migrate:
	goose postgres ./migrations create add_some_column sql

local-migrate:
	&(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v



vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
#		@if [ ! -d vendor.protogen/google ]; then \
#			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
#			mkdir -p  vendor.protogen/google/ &&\
#			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
#			rm -rf vendor.protogen/googleapis ;\
#		fi
#		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
#			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
#			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
#			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
#			rm -rf vendor.protogen/openapiv2 ;\
#		fi


gen-cert:

	openssl genrsa -out keys/ca.key 4096
	openssl req -new -x509 -key keys/ca.key -sha256 -subj '/C=US/ST=NJ/O=CA, Inc.' -days 365 -out keys/ca.cert
	openssl genrsa -out keys/service.key 4096
	openssl req -new -key keys/service.key -out keys/service.csr -config keys/certificate.conf
	openssl x509 -req -in keys/service.csr -CA keys/ca.cert -CAkey keys/ca.key -CAcreateserial \
    		-out keys/service.pem -days 365 -sha256 -extfile keys/certificate.conf -extensions req_ext