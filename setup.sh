#!/bin/bash

set -euxo pipefail

# $1 --> project name 
# $2 --> package name

git clone git@github.com:typical-go/rest-project-starter.git $1
rm -rf $1/.git
rm -rf $1/setup.sh $1/README.md
find $1 -type f -exec sed -i '' "s+github.com/typical-go/rest-project-starter+$2+g" {} \;
find $1 -type f -exec sed -i '' "s+rest-project-starter+$1+g" {} \;
mv $1/cmd/rest-project-starter $1/cmd/$1