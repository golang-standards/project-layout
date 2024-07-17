# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest as builder
RUN go version

# Copy the local package files to the container's workspace.
ADD . /go/src/gitlab.com/{OWNER_NAME}/{PACKAGE_NAME}

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/gitlab.com/{OWNER_NAME}/{PACKAGE_NAME}

# Build the output command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN set -x && \
    go get github.com/golang/dep/cmd/dep && \
    dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o cmd/{PACKAGE_NAME} $GOPATH/src/gitlab.com/{OWNER_NAME}/{PACKAGE_NAME}/cmd

# Stage 2 (to create a downsized "container executable", ~7MB)

# If you need SSL certificates for HTTPS, replace `FROM SCRATCH` with:
#
#   FROM alpine:3.7
#   RUN apk --no-cache add ca-certificates
#
FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/gitlab.com/{OWNER_NAME}/{PACKAGE_NAME}/cmd/{PACKAGE_NAME} .
#COPY --from=builder /go/src/gitlab.com/{OWNER_NAME}/{PACKAGE_NAME}/configs/json/db.json /configs/json/db.json
#COPY --from=builder /go/src/gitlab.com/{OWNER_NAME}/{PACKAGE_NAME}/configs/json/server.json /configs/json/server.json
#RUN ls
# Run the output command by default when the container starts.
CMD ["./{PACKAGE_NAME}"]

# Document that the service listens on port 8080.
EXPOSE 8080

