#!/bin/bash

# Select Servic name
SERVICE_PATH=$1
SERVICE_NAME=$2
RELEASE_VERSION=$3
CONTAINER_REGISTRY=$4

echo service path: $SERVICE_PATH
echo service name: $SERVICE_NAME
echo service version: $RELEASE_VERSION
echo container registry: $CONTAINER_REGISTRY

# Create packages
mkdir -p packages/$SERVICE_PATH

# COPY
cp $SERVICE_PATH/deployments/* packages/$SERVICE_PATH

# UPDATE manifest.json
sed -i '' 's/RELEASE_VERSION/'"$RELEASE_VERSION"'/' packages/$SERVICE_PATH/manifest.json

# docker save to .tgz
docker save $CONTAINER_REGISTRY/$SERVICE_NAME | gzip > packages/$SERVICE_PATH/$SERVICE_NAME.tgz

# create package
tar zcfv packages/$SERVICE_PATH/$SERVICE_NAME.spx *
