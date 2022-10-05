#!/usr/bin/env bash
set -e pipefail

## Make sure all required tools are installed
# shellcheck disable=SC2039
declare -A tools
tools=(
    [migrate]="migrate --version"
    [sqlc]="sqlc version"
    [go]="go version"
    [docker]="docker version"
    [docker-compose]="docker-compose version"
    [kubectl]="kubectl version"
    [helm]="helm version"
)

for tool in "${!tools[@]}"; do
    if ! command -v "${tool}" >/dev/null 2>&1; then
        echo "ERROR: ${tool} is not installed"
        exit 1
    fi
    if ! ${tools[${tool}]} >/dev/null 2>&1; then
        echo "ERROR: ${tool} is not working properly"
        exit 1
    fi
done

## Create directory structure for the project
# db/migrations -- contains all db migrations
# db/query -- contains all sql queries for sqlc
# db/sqlc -- contains all generated code from sqlc (do not edit)
# db/mock -- contains mock functions for the database
mkdir -p db/{migrations,query,sqlc,mock}

## Create sqlc file for the project

cat <<EOF > ./sqlc.yaml
# auto-generate by ./hack/init.sh (DO NOT EDIT)
version: "2"
sql:
  - schema: "db/migrations"
    queries: "db/query"
    engine: "postgresql"
    gen:
      go:
        package: "database"
        out: "db/sqlc"
        emit_interface: true
EOF