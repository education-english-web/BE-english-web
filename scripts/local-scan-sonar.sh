#!/bin/bash
empty=
FILES=
if [ -z "$id" ]
then
    FILES=`git status | grep 'new file\|renamed\|modified'`
    FILES=`echo $FILES | sed 's/new file: / /g; s/renamed: / /g; s/modified: / /g'`
    FILES=`echo $FILES | sed -E 's|.* -> | |'`
    FILES=`echo $FILES | xargs -n1 | sort -u`
    FILES=`echo $FILES | sed 's/ /,/g'`
else
    FILES=`git diff-tree --no-commit-id --name-status ${id} -r`
    FILES=`echo $FILES | sed 's/M / /g; s/A / /g'`
    FILES=`echo $FILES | sed -E 's|D .* | |'`
    FILES=`echo $FILES | sed 's/ /,/g'`
fi

sonar-scanner -Dsonar.projectKey=stampless_backend \
    -Dsonar.sources=${FILES} \
    -Dsonar.host.url=${SONAR_HOST_URL} \
    -Dsonar.login=${SONAR_QUBE_TOKEN}
