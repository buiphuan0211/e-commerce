#!bin/bash
export ENV=develop
export DOMAIN_ECOMMERCE_ADMIN=localhost:9000
export ZOOKEEPER_URI=127.0.0.1:2181
export ZOOKEEPER_PREFIX_ECOMMERCE_COMMON=/ecommerce/common
export ZOOKEEPER_PREFIX_ECOMMERCE_ADMIN=/ecommerce/admin



run-admin:
	go run cmd/admin/main.go


swagger-admin:
	swag init -d ./ -g cmd/admin/main.go \
    -o ./docs/admin