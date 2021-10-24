FROM envoyproxy/envoy-dev:4b5b3386c6b0d2d284bb1f71639c8e0972659867
COPY envoy.yaml /etc/envoy/envoy.yaml
RUN chmod go+r /etc/envoy/envoy.yaml