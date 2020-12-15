#!/bin/bash

set -e

SRC_CODE_DIR="${SRC_CODE_DIR:-defaultdirectory}"

main(){
    # Change directory to the code
    echo "This is SRC CODE: ${SRC_CODE_DIR}"
	pushd "${SRC_CODE_DIR}"
        # Execute the Unit tests and produce coverage report
        go test -coverprofile=coverage.out
        # Parse and present the coverage report as part of the logs
        go tool cover -func=coverage.out
    popd
}

main
