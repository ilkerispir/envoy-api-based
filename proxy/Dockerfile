FROM envoyproxy/envoy:v1.20.0

COPY ./envoy.yaml /etc/envoy.yaml

RUN chmod go+r /etc/envoy.yaml

CMD ["/usr/local/bin/envoy", "-c /etc/envoy.yaml", "-l", "debug"]