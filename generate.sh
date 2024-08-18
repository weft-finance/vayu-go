#!/bin/bash

# Directory path
directory_path="./open-api-files"

# File name within the directory
file_name="open-api.yaml"

# Check if the directory exists
if [[ ! -d "$directory_path" ]]; then
    echo "Directory '$directory_path' does not exist."
    exit 1
fi

# Check if the file exists within the directory
if [[ ! -f "$directory_path/$file_name" ]]; then
    echo "File '$file_name' does not exist in directory '$directory_path'."
    exit 1
fi

if [[ -f "openapi" ]]; then
    rm -rf openapi
fi


echo "Generating go client for openapi"

npx @openapitools/openapi-generator-cli generate -g go --additional-properties=prependFormOrBodyParameters=true -o openapi -i ./open-api-files/open-api.yaml

rm -rf openapi/go.mod
rm -rf openapi/go.sum
rm -rf openapi/test

echo "Generated go client for openapi"