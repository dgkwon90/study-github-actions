# sutdy-github-actions

## Command 정리

### Docker build

root diretory를 기준으로 아래의 command를 입력한다.

``` shell
# server
docker build --build-arg SERVICE_NAME=server --build-arg SERVICE_PORT=9999 --build-arg GIT_SHORT_COMMIT_ID=$(git rev-parse --short HEAD) --build-arg BUILD_TIME=$(date +%Y-%m-%d_%H:%M) -t server:1.0.0 -f server/build/Dockerfile .

# client
docker build --build-arg SERVICE_NAME=client --build-arg GIT_SHORT_COMMIT_ID=$(git rev-parse --short HEAD) --build-arg BUILD_TIME=$(date +%Y-%m-%d_%H:%M) -t client:1.0.0 -f client/build/Dockerfile .
```

### Docker-compose up

root diretory를 기준으로 아래의 command를 입력한다.

```shell
# server
docker-compose -f server/deployments/docker-compose.yaml up -d

#client
docker-compose -f client/deployments/docker-compose.yaml up -d
```
