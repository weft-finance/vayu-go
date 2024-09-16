#!/bin/bash

version=$1

if [ -z "$version" ]; then
  echo "No version provided. Usage: add_supported_version.sh <version>"
  exit 1
fi

pushd "$(pwd)" > /dev/null

./version_support/validate_go_version.sh "$version"
if [ $? -ne 0 ]; then
  echo "Version $version is not valid."

  popd > /dev/null
  exit 1
fi

if grep -Fxq "$version" ./version_support/supported_versions.txt; then
  echo "Version $version is already supported."
else
  echo -e "\n$version" >> ./version_support/supported_versions.txt
  echo "Version $version has been added to supported_versions.txt."
fi

popd > /dev/null