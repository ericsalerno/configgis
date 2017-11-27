FROM golang

# Copy local package to the containers workspace
ADD . /go/src/github.com/ericsalerno/configgis

# Build package in container
RUN go install github.com/ericsalerno/configgis

WORKDIR "/go/bin"

# Set container entrypoint to compiled binary
ENTRYPOINT configgis

ENV CFIGGIS_PORT 765

# Expose port
EXPOSE ${CFIGGIS_PORT}