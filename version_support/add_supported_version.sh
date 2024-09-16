#!/bin/bash

version=$1

if [ -z "$version" ]; then
  echo "No version provided. Usage: add_supported_version.sh <version>"
  exit 1
fi

pushd "$(pwd)" > /dev/null

if grep -Fxq "$version" ./version_support/supported_versions.txt; then
  echo "Version $version is already supported."
  exit 0
fi

./version_support/validate_go_version.sh "$version"
if [ $? -ne 0 ]; then
  echo "Version $version is not valid."

  popd > /dev/null
  exit 1
fi


if [ "$SET_LATEST" == "true" ]; then
  go mod edit -go="$version"
fi

echo -e "\n$version" >> ./version_support/supported_versions.txt
echo "Version $version has been added to supported_versions.txt."

popd > /dev/null