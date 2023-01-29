FROM scratch
COPY feserve /usr/bin/feserve

ENTRYPOINT ["/usr/bin/feserve"]
