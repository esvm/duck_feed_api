FROM alpine:latest
RUN apk add --update ca-certificates
COPY bin/api_gateway /
RUN chmod +x api_gateway
ENTRYPOINT ["/api_gateway"]