#!/bin/bash

# This is intended to be invoked by the Makefile.

declare -a PACKAGES_PATHS=(
)

find ./internal -type f -name "*_mock.go" -print0 | xargs -0 rm

for PACKAGES_PATH in "${PACKAGES_PATHS[@]}"
do
  PACKAGE=$(echo "${PACKAGES_PATH}" | sed -nE 's/.*\/([^/]*)$/\1/p' )
  mockgen -source="${PACKAGES_PATH}/interfaces_mockable.go" -package=${PACKAGE} -destination "${PACKAGES_PATH}/interfaces_mock.go"
done