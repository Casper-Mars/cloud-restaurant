API_PROTO_FILES=$(shell find api -name *.proto)


.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
               --validate_out=paths=source_relative,lang=go:. \
               --openapiv2_out . \
	       $(API_PROTO_FILES)

.PHONY: set-up-dev
set-up-dev:
	docker-compose -f ./hack/env/docker-compose.yml up -d

.PHONY: stop-dev
stop-dev:
	docker-compose -f ./hack/env/docker-compose.yml down
