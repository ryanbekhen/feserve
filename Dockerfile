FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY feserve /usr/bin/feserve

ENTRYPOINT ["/usr/bin/feserve"]
