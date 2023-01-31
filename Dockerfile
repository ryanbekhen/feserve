FROM alpine
RUN apk update \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates
COPY feserve /usr/bin/feserve

ENTRYPOINT ["/usr/bin/feserve"]
