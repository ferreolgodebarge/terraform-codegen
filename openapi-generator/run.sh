#!/bin/bash

java -jar openapi-generator-cli.jar generate -l go -i http://127.0.0.1:5000/api/v1/swagger.json -o src/openapi
