# TODO: Add Docker tasks here.
FROM supinf/golangci-lint:latest AS build-env
WORKDIR /go/src/github.com/miraikeitai2020/backend-bff
COPY ./ ./
# Install gcc
RUN apk --update add libc-dev
RUN apk --update add openssl
# generate private key
RUN openssl genrsa -out private.key 4096
# setup for build API
RUN go get github.com/99designs/gqlgen
RUN go run github.com/99designs/gqlgen
# build API
RUN go build -o server main.go

FROM alpine:latest
WORKDIR /usr/local/bin/
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/miraikeitai2020/backend-bff/server /usr/local/bin/server
COPY --from=build-env /go/src/github.com/miraikeitai2020/backend-bff/private.key /usr/local/bin/private.key
ENV PORT 9020

EXPOSE 9020
CMD ["/usr/local/bin/server"]
