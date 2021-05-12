FROM golang:latest AS build-env

# allow Go to retrive the dependencies for the build step
RUN apk add --no-cache git

# secure against running as root
RUN adduser -D -u 10000 goUser
RUN mkdir /service/ && chown goUser /service/
USER goUser

WORKDIR /service/
ADD . /service/

# compile the binary, we don't want to run the cgo resolver
RUN CGO_ENABLED=0 go build -o /service/ .

# final stage
FROM alpine:latest

# secure against running as root
RUN adduser -D -u 10000 goUser
USER goUser

WORKDIR /
COPY --from=build-env /service/certs/docker.localhost.* /
COPY --from=build-env /service/ /

EXPOSE 8080

CMD ["/"]
