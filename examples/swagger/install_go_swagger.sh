#!/bin/bash

if [[ $VERSION == "" ]]; then
  VERSION="v0.27.0"
fi

# Ensure the bin directory exists
mkdir -p ./bin

# Check the current version
if [[ -f "./bin/swagger" ]]; then
  CURRENT_VERSION=$(./bin/swagger version | head -n 1 | cut -d ":" -f2 | xargs)
  if [[ "$CURRENT_VERSION" == "$VERSION" ]]; then
    echo "go-swagger $VERSION already downloaded"
    exit 0
  else
    echo "found version $CURRENT_VERSION, need version $VERSION"
    rm ./bin/swagger
  fi
fi

# Download Go Swagger if required
if [[ ! -f "./bin/swagger" ]]; then
  if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "downloading go-swagger $VERSION for mac"
    curl -L "https://github.com/go-swagger/go-swagger/releases/download/$VERSION/swagger_darwin_amd64" -o ./bin/swagger
  else
    echo "downloading go-swagger $VERSION for linux"
    wget "https://github.com/go-swagger/go-swagger/releases/download/$VERSION/swagger_linux_amd64" -O ./bin/swagger
  fi
  chmod +x ./bin/swagger
  echo "go-swagger $VERSION installed"
fi
