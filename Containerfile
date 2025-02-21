FROM index.docker.io/library/golang:alpine as build
WORKDIR /src
COPY main.go go.mod go.sum .
RUN go build -o demo -ldflags="-s -w"

FROM index.docker.io/library/alpine:edge
LABEL maintainer "Raimond van Stijn <ramons@nl.ibm.com>"
RUN addgroup -g 1970 demo \
    && adduser -u 1970 -G demo -s /bin/sh -D demo
COPY --chown=demo:demo --from=build /src/demo /app/demo
USER demo
EXPOSE 8080
ENTRYPOINT /app/demo
