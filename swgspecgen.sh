#! /usr/bin/env bash

dir=`dirname $0`

echo -e "[url \"git@gitlab.wallstcn.com:\"]\n\tinsteadOf = https://gitlab.wallstcn.com/" >> ~/.gitconfig

go get -u gitlab.wallstcn.com/opensource/goswgspecgen

# git clone git@gitlab.wallstcn.com:opensource/goswgspecgen.git /opt/go/src/gitlab.wallstcn.com/opensource/oswgspecgen

goswgspecgen \
-apiPackage="github.com/songrgg/backeye/service/api" \
-mainApiFile="$dir/main.go" \
-output="$dir/public/swagger-resources" \
-format="swagger"

#sed -i '/"basePath"/d' $dir/public/swagger-resources/index.json
