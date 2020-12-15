#!/bin/bash

set -e

SRC_CODE_DIR="${SRC_CODE_DIR:-defaultdirectory}"

main(){
    # Change directory to the code
    echo "This is SRC CODE: ${SRC_CODE_DIR}"
	pushd "${SRC_CODE_DIR}"
        # Download and Install the Golint
        go get -u golang.org/x/lint/golint
        # Run Golint and return non-zero exit code if there are lint issues
        golint -set_exit_status
    popd
}

main
