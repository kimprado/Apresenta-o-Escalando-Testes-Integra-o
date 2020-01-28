#!/bin/sh

set -e

case "$1" in
    build)
        go build \
        -o ./equipment-rental-api.bin \
        github.com/example/equipment-rental/cmd/equipmentRentalAPI 
        ;;
    wire)
        wire ./cmd/equipmentRentalAPI/
        ;;
    wire-testes)
        wire gen --output_file_prefix=tmp_ \
        ./internal/pkg/commom/config \
        ./internal/pkg/infra/mongo \
        ./internal/pkg/infra/redis 

        mv ./internal/pkg/commom/config/tmp_wire_gen.go ./internal/pkg/commom/config/wire_gen_test.go
        mv ./internal/pkg/infra/mongo/tmp_wire_gen.go ./internal/pkg/infra/mongo/wire_gen_test.go
        mv ./internal/pkg/infra/redis/tmp_wire_gen.go ./internal/pkg/infra/redis/wire_gen_test.go
        ;;
    *)
        echo "Usage: {build|wire|wire-testes}" >&2
       ;;
esac
