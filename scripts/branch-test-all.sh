#!/bin/bash

TEST_RESULT_DIR="${TEST_RESULTS:-./test-results}"
CHROME_PATH="${CHROME_PATH_ENV:-/Applications/Google Chrome.app/Contents/MacOS/Google Chrome}"

mkdir -p ${TEST_RESULT_DIR}
echo "Listing all packages"
PKG_LIST+="$(go list ./... | grep -v vendor/ | grep -v scripts | grep -v migrations | grep -v mock | grep -v datapatch | grep -v db/ | grep -v app/domain/repository ) "

echo "----------------"
echo "Branch Test:"

NEW_LIST=()
ROOT_PATH="github.com/education-english-web/BE-english-web/"
>${TEST_RESULT_DIR}/gobco_out.log
>${TEST_RESULT_DIR}/branch-cover.json

for pack in ${PKG_LIST}
do
    new_pack=${pack#"$ROOT_PATH"}
    NEW_LIST+="${new_pack} "
done

for p in ${NEW_LIST}
do
    echo "Analysing: ${p}"
    gobco -branch -list-all -test -v -stats ${p}/branch-cover.json ${p} >>${TEST_RESULT_DIR}/gobco_out.log
    if [ -f ${p}/branch-cover.json ]; then
        go-branch-coverage -srcPackageName ${p} -targetPackageName test-results -isUpdateJson=true
        rm ${p}/branch-cover.json
    fi
done

go-branch-coverage -srcPackageName ${TEST_RESULT_DIR} -isUpdateJson=false
