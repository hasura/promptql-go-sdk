#!/bin/bash

docker run --rm -v ./:/local openapitools/openapi-generator-cli:latest generate \
    -i https://api.promptql.pro.hasura.io/openapi.json \
    -g go \
    -t /local/template \
    -c /local/openapi.config.json \
    -o /local/api

SED_CMD="sed -i"
if [ "$(uname)" == "Darwin" ]; then
  SED_CMD="sed -i .bak"
fi

API_DEFAULT_FILE="api/api_default.go"
# patch some custom types because of the limitation of the generation tool
$SED_CMD 's/request \*Request/request *ExecuteRequest/g' $API_DEFAULT_FILE
$SED_CMD 's/request Request/request ExecuteRequest/g' $API_DEFAULT_FILE
$SED_CMD 's/request1 \*Request1/request1 *QueryRequest/g' $API_DEFAULT_FILE
$SED_CMD 's/request1 Request1/request1 QueryRequest/g' $API_DEFAULT_FILE
$SED_CMD 's/\*map\[string\]interface{}/map[string]interface{}/g' "api/model_table_artifact.go"

rm -f api/*.go.bak