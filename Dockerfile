FROM scratch
COPY feserve /usr/bin/feserve
EXPOSE 8000

ENTRYPOINT ["/usr/bin/feserve"]
