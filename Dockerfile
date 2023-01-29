FROM scratch
COPY feserve /usr/bin/feserve

EXPOSE 80
EXPOSE 443

ENTRYPOINT ["/usr/bin/feserve"]
