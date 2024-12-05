#!/bin/bash

set -eu
set -o pipefail

TEST_RESULT_DIR="${TEST_RESULTS:-./test-results}"
mkdir -p ${TEST_RESULT_DIR}

PKG_LIST+="$(go list ./... | grep -v /vendor/ | grep -v migrations) "
echo "----------------"

echo "Running test"
go test -covermode=count ${PKG_LIST} -coverprofile=${TEST_RESULT_DIR}/coverage.out -json > ${TEST_RESULT_DIR}/test.out
