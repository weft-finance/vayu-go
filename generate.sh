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

yes | npx @openapitools/openapi-generator-cli generate -g go --additional-properties=prependFormOrBodyParameters=true,disallowAdditionalPropertiesIfNotPresent=false,enumClassPrefix=true,isGoSubmodule=true,withGoMod=false -o openapi -i ./open-api-files/open-api.yaml
rm -rf openapi/test
rm -rf openapi/README.md
rm -rf openapi/git_push.sh
rm -rf openapi/.travis
rm -rf openapi/.openapi-generator
rm -rf openapi/docs

echo "Generated go client for openapi"


# # Directory where the generated code is located
# DIR="openapi"

# # Find all Go files in the directory
# find "$DIR" -name "*.go" | while read -r file; do
#     # Use sed to match and replace constant names with the prefixed type
#     sed -i '' -E 's/([A-Z_]+) ([A-Za-z]+) = ("[A-Za-z_]+")/\2_\1 \2 = \3/g' "$file"
# done

# echo "Post-processing completed. Constants are prefixed with their type."

# # fix2 

# # Define the path to the specific file
# FILE="openapi/model_plan_billing_data.go"

# # Use sed to find and replace POSTPAYMENT with PaymentTerm_POSTPAYMENT
# sed -i '' 's/POSTPAYMENT/PaymentTerm_POSTPAYMENT/g' "$FILE"

# echo "Replacement completed: POSTPAYMENT -> PaymentTerm_POSTPAYMENT in $FILE"

