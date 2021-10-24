## Envoyproxy API-based dynamic endpoint discovery configuration EDS, LDS, CDS

## Architecture
![Architecture](architecture.png)

## Run Envoy
```
func-e run -c envoy/envoy.yaml

docker run -d --name envoy -e ENVOY_UID=777 -e ENVOY_GID=777 -p 9901:9901 -p 80:10000 -v /root/:/etc/envoy ilkerispir/envoy

docker run -d --name envoy -e ENVOY_UID=777 -e ENVOY_GID=777 -p 9901:9901 -p 80:10000 envoyproxy/envoy-dev:4b5b3386c6b0d2d284bb1f71639c8e0972659867
```
## Run Upstream
```
docker run -p 8081:8081 -d ilkerispir/upstream
```

## Run xDS
```
docker run -p 8080:8080 -d ilkerispir/xds
```

## Remove all containers & images
```
docker rm -vf $(docker ps -a -q)
docker rmi -f $(docker images -a -q)
```

https://hub.docker.com/u/ilkerispir