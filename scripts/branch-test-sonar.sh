#!/bin/bash

TEST_RESULT_DIR="${TEST_RESULTS:-./test-results}"

mkdir -p ${TEST_RESULT_DIR}

echo "Listing Changed Packages in PR..."
PACK_LIST_IN_PR+=$(cat ${TEST_RESULT_DIR}/changed-package.txt | sort -u)

PKG_LIST+=($(tr ' ' '\n' <<<"${PACK_LIST_IN_PR[@]}"))
echo "----------------"
echo "Branch Test:"

NEW_LIST=()
ROOT_PATH="github.com/education-english-web/BE-english-web/"

>${TEST_RESULT_DIR}/gobco_out.log
>${TEST_RESULT_DIR}/branch-cover.json

for pack in "${PKG_LIST[@]}"; do
    if [[ ${pack:0:1} == '.' ]]; then
        continue
    fi
    new_pack=${pack#"$ROOT_PATH"}
    NEW_LIST+="${new_pack} "
done

for p in ${NEW_LIST}
do
    echo "Analysing: ${p}"
    gobco -branch -list-all -test -v -stats ${p}/branch-cover.json ${p} >>${TEST_RESULT_DIR}/gobco_out.log
    echo "Analysing brnach coverage for: ${p}" >>${TEST_RESULT_DIR}/gobco_out.log
    if [ -f ${p}/branch-cover.json ]; then
        go-branch-coverage -srcPackageName ${p} -targetPackageName test-results -isUpdateJson=true
        rm ${p}/branch-cover.json
    fi
done

go-branch-coverage -srcPackageName ${TEST_RESULT_DIR} -isUpdateJson=false
