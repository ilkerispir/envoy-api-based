## Envoyproxy API-based dynamic endpoint discovery configuration

## Architecture
![Architecture](images/architecture.png)

## Run Envoyproxy
```
func-e run -c envoy.yaml
```

## Run xDS & Resource
```
docker-compose up -d
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