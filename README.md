# docker-grpc-service-tutorial

A tutorial that shows how to dockerize a gRPC service written in Golang.

## running it

```
make run
```

## using Docker

### regular image build

```
docker build . -t docker-grpc-service-tutorial
```

### smaller image build

```
docker build -f Dockerfile.multistage  . -t docker-grpc-service-tutorial
```

### running it

```
docker run -p 4040:4040 docker-grpc-service-tutorial
```
