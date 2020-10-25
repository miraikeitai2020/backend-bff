GOCMD=go
DOCKERCMD=docker
SSL=openssl genrsa
GO_RUN=$(GOCMD) run
GO_BUILD=$(GOCMD) build
DOCKER_BUILD=$(DOCKERCMD) build
DOCKER_RUN=$(DOCKERCMD) run

GQLGEN=github.com/99designs/gqlgen

all:
	$(SSL) -out private.key 4096
	$(GO_RUN) $(GQLGEN)
clean:
	rm private.key
	rm -rf pkg/bff
	rm -rf pkg/server/model
build:
	$(GO_BUILD) -o server main.go
run:
	$(GO_RUN) main.go
docker-build:
	$(DOCKER_BUILD) ./ -t miraikeitai2020/bff:0.2.0
docker-run:
	$(DOCKER_RUN) -d -p 9020:9020 miraikeitai2020/bff:0.2.0
