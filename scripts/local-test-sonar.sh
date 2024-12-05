#!/bin/bash

TEST_RESULT_DIR="${TEST_RESULTS:-./test-results}"

# statement test
if [[ $1 =~ ^"all"|"statement"$ ]]
then
    echo "Statement Test"
    PKG_LIST+="$(go list ./... | grep -v /vendor/ | grep -v migrations) "
    go test -v -covermode=count ${PKG_LIST} -coverprofile=${TEST_RESULT_DIR}/coverage.out
fi

if ! [[ $1 =~ ^"all"|"branch"$ ]]
then
    exit 0
fi

# branch test
NEW_LIST=()
PKG_LIST=()
FILES=
ROOT_PATH="github.com/education-english-web/BE-english-web/"

# prepare test result file
>./test-results/gobco_out.log
>./test-results/branch-cover.json

echo "Branch Test"

# prepare changed files
if [ -z "$id" ]
then
    FILES=`git status | grep 'new file\|renamed\|modified'`
    FILES=`echo $FILES | sed 's/new file: / /g; s/renamed: / /g; s/modified: / /g'`
    FILES=`echo $FILES | sed -E 's|.* -> | |'`
    FILES=`echo $FILES | xargs -n1 | sort -u`
else
    FILES=`git diff-tree --no-commit-id --name-status ${id} -r`
    FILES=`echo $FILES | sed 's/M / /g; s/A / /g'`
    FILES=`echo $FILES | sed -E 's|D .* | |'`
fi

# extract to package name only
for file in "${FILES[@]}"; do
    PKG_LIST+=`dirname $file | awk -F "/[^/]/*$" '{ print ($1 == "." ? "": $1); }' | sort | uniq`
done

for pack in "${PKG_LIST[@]}"; do
    if [[ ${pack:0:1} == '.' ]]; then
        continue
    fi
    new_pack=${pack#"$ROOT_PATH"} # trim prefix ROOT_PATH
    NEW_LIST+="${new_pack} "
done

# test branch coverage for each package
for p in ${NEW_LIST}
do
    echo "Analysing: ${p}"
    gobco -branch -list-all -test -v -stats ${p}/branch-cover.json ${p} >>${TEST_RESULT_DIR}/gobco_out.log
    if [ -f ${p}/branch-cover.json ]; then
        go-branch-coverage -srcPackageName ${p} -targetPackageName test-results -isUpdateJson=true
        rm ${p}/branch-cover.json
    fi
done

# generate branch coverage result
go-branch-coverage -srcPackageName ${TEST_RESULT_DIR} -isUpdateJson=false
