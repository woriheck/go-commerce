#!/bin/bash

# Retrieve code coverage from infra/github/.out/coverage.txt
BASEDIR=$(dirname "$0")
mkdir -p $BASEDIR/.out
coverage=$(go tool cover -func $BASEDIR/.out/coverage.log | grep total | awk '{print $3}')
num=$(echo $coverage | sed "s/%//g")
num=$(printf "%.0f" $num)

# Code coverage
#   >= 80% - Green
#   >= 60% - Yellow
#   <  60% - Red

if [[ "$num" -gt 80 ]]
then
    color=green
elif [[ "$num" -gt 69 ]]
then
    color=yellow
else
    color=red
fi

# Download code coverage badge
curl https://badgen.net/badge/coverage/"$coverage"25/$color --retry 2 -o $BASEDIR/.out/coverage.svg
