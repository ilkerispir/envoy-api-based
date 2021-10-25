## Envoyproxy API-based dynamic endpoint discovery configuration

## Architecture
![Architecture](images/architecture.png)

## Run Envoy
```
cd envoy; func-e run -c envoy.yaml
```

## Run xDS
```
cd xds; go run grpc_server.go
```

## Run Resource
```
cd resource; python server.py -p 8081
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