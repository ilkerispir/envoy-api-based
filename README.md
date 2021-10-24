## Envoyproxy API-based dynamic endpoint discovery configuration EDS, LDS, CDS

## Architecture
![Architecture](architecture.png)

## Commands
```
docker run -d --name envoy -p 9901:9901 -p 80:10000  -v /root/:/etc/envoy ilkerispir/envoy
```

```
docker run -p 8081:8081 -d ilkerispir/upstream
```

```
docker run -p 8080:8080 -d ilkerispir/xds
```

## Remove all containers & images
```
docker rm -vf $(docker ps -a -q)
docker rmi -f $(docker images -a -q)
```

https://hub.docker.com/u/ilkerispir