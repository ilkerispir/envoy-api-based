## Envoyproxy API-based dynamic endpoint discovery configuration

## Architecture
![Architecture](images/architecture.png)

## Run Envoy
```
docker run -d -p 80:80 -p 9000:9000 ilkerispir/envoy
```

## Run xDS
```
cd xds; go run grpc_server.go
```

## Run Resource
```
docker run -d -p 8081:8080 ilkerispir/resource
```

## cURL Test
```
while true; do curl http://localhost:10000; sleep .5; printf '\n'; done
```

## Result
![Result](images/curl.png)

## Remove all containers & images
```
docker rm -vf $(docker ps -a -q)
docker rmi -f $(docker images -a -q)
```

https://hub.docker.com/u/ilkerispir