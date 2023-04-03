#!/bin/bash

# Select Servic name
SERVICE_PATH=$1
SERVICE_NAME=$2
RELEASE_VERSION=$3

echo service path: $SERVICE_PATH
echo service name: $SERVICE_NAME
echo service version: $RELEASE_VERSION

# Create packages
mkdir -p packages/$SERVICE_PATH

# COPY
cp $SERVICE_PATH/deployments/* packages/$SERVICE_PATH

# UPDATE manifest.json
sed -i '' 's/RELEASE_VERSION/'"$RELEASE_VERSION"'/' packages/$SERVICE_PATH/manifest.json

# docker save
docker save $SERVICE_NAME | gzip > packages/$SERVICE_PATH/$SERVICE_NAME.tgz

# create package
tar zcfv packages/$SERVICE_PATH/$SERVICE_NAME.spx *
