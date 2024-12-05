#!/bin/bash

RED=`tput setaf 1`
GRN=`tput setaf 2`
PUR=`tput setaf 13`
RESET=`tput sgr0`
TEST_RESULT_DIR="${TEST_RESULTS:-./test-results}"
mkdir -p ${TEST_RESULT_DIR}
echo "${GRN}Listing all packages${RESET}"
PKG_LIST+="$(go list ./... | grep -v /vendor/ | grep -v /mock | grep -v migrations) "
for i in $PKG_LIST
do
    echo $i
done
echo "----------------"
EXIT_CODE=0
echo "${GRN}Golangci-lint...${RESET}"
golangci-lint --timeout 15m0s run -v ./...
EXIT_CODE+=$?
if [ $EXIT_CODE -eq 0 ]
then
    echo "${GRN}Success${RESET}"
else
    echo "${RED}Failed${RESET}"
fi
exit $EXIT_CODE 
