###########################################
# build
###########################################
# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.20-alpine AS build

# Docker build ARG
ARG SERVICE_NAME=server
ARG SERVICE_PORT=9999
ARG GIT_SHORT_COMMIT_ID=$(git rev-parse --short HEAD)
ARG BUILD_TIME=$(date +'%Y-%m-%d_%H:%M')

###########################################
# Git package installation.
RUN apk add --no-cache git

###########################################
# Defining Go Build Options
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

    # Create and change to the app directory.
WORKDIR /app

COPY . ./

# Build the binary.
RUN go build -v -o service \
    -ldflags "-X 'main.GitCommit=${GIT_SHORT_COMMIT_ID}' -X 'main.BuildTime=${BUILD_TIME}'" \
    ./${SERVICE_NAME}/cmd/${SERVICE_NAME}/main.go

###########################################
# Base
############################################
# This image is most useful in the context of building base images (such as debian and busybox)
# or super minimal images
# (that contain only a single binary and whatever it requires, such as hello-world).
# https://hub.docker.com/_/scratch
FROM scratch

# Copy the binary to the production image from the builder stage.
COPY --from=build /app/service service

EXPOSE ${SERVICE_PORT}

# Run the service on container startup.
ENTRYPOINT [ "/service" ]
