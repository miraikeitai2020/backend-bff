FROM supinf/golangci-lint:latest AS build-env
WORKDIR /go/src/github.com/miraikeitai2020/backend-bff
COPY ./ ./
# Install gcc
RUN apk --update add libc-dev
# setup for build API
RUN rm /go/src/github.com/miraikeitai2020/backend-bff/pkg/server/resolver/production.go
RUN mv /go/src/github.com/miraikeitai2020/backend-bff/config/development/gqlgen.yml /go/src/github.com/miraikeitai2020/backend-bff/gqlgen.yml 
RUN go get github.com/99designs/gqlgen
RUN go run github.com/99designs/gqlgen
# build API
RUN go build -o server pkg/server/server.go

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /go/src/github.com/miraikeitai2020/backend-bff/server /usr/local/bin/server
ENV PORT 9020

EXPOSE 9020
CMD ["/usr/local/bin/server"]