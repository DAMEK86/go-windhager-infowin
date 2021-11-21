
#!/usr/bin/env bash

red='\033[1;31m'
green='\033[1;32m'
normal='\033[0m'

appName='windhager-infowin-exporter'

export CGO_ENABLED=0
goflags=""
if [[ "${READ_ONLY:-false}" == "true" ]]; then
    echo "Running in readonly mode"
    goflags="-mod=readonly"
fi

## go-fmt: format go code
function task_go_fmt {
    go fmt ./...
}

## build-cli [OS]: builds the go executable for cli
function task_build_cli {
  GOOS=$1 go build -a ${goflags} -ldflags="-s -w" -o ${appName}-cli cmd/main.go
}

## build-influx [OS]: builds the go executable with influx for container usage
function task_build_influx {
  GOOS=$1 go build -a -o ${appName} -trimpath -ldflags="-s -w" ${goflags} internal/main.go
}

## build-container: builds the container image
function task_build_container {
    task_build_influx "linux"
    docker build -t damek/${appName} .
}

function task_usage {
    echo "Usage: $0"
    sed -n 's/^##//p' <$0 | column -t -s ':' |  sed -E $'s/^/\t/'
}

CMD=${1:-}
shift || true
RESOLVED_COMMAND=$(echo "task_"$CMD | sed 's/-/_/g')
if [ "$(LC_ALL=C type -t $RESOLVED_COMMAND)" == "function" ]; then
    pushd $(dirname "${BASH_SOURCE[0]}") >/dev/null
    $RESOLVED_COMMAND "$@"
else
    task_usage
fi