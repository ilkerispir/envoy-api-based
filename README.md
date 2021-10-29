## Envoyproxy API-based dynamic endpoint discovery configuration

## Architecture
![Architecture](images/architecture.png)

## Run Envoy
```
func-e run -c envoy.yaml

docker run -d -p 10000:10000 -p 9000:9000 ilkerispir/envoy
```

## Run xDS
```
cd xds; go run grpc_server.go
docker run -d -p 8080:8080 -p 5000:5000 ilkerispir/xds
```

## Run Resource
```
docker run -d -p 8081:8081 ilkerispir/resource
```

## Add Resource
```
curl http://localhost:5000/edsservice/register?endpoint=127.0.0.1:8081
```

## cURL Test
```
while true; do curl http://localhost:10000; sleep .5; printf '\n'; done
```

## Result
![Result](images/curl.png)

## Start all of our containers
```
docker-compose up -d
```

## Remove all containers & images
```
docker rm -vf $(docker ps -a -q)
docker rmi -f $(docker images -a -q)
```

https://hub.docker.com/u/ilkerispir