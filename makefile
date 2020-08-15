GOCMD=go
DOCKERCMD=docker
GORUN=$(GOCMD) run
GOGET=$(GOCMD) get
BUILDIMG=$(DOCKERCMD) build
RUNIMG=$(DOCKERCMD) run
TMPGEN= mkdir tmp
DEVDOCKER=config/development/Dockerfile
DEVCONFIG=config/development/gqlgen.yml
GQLGEN=github.com/99designs/gqlgen

all:
	$(GOGET) $(GQLGEN)
clean:
	mv tmp/*.go pkg/server/resolver/
	rm -r tmp
	rm *.yml
	rm -rf pkg/bff pkg/server/model
local-dev-run:
	$(TMPGEN)
	mv pkg/server/resolver/production.go tmp/
	cp $(DEVCONFIG) gqlgen.yml
	$(GORUN) $(GQLGEN)
	$(GORUN) pkg/server/server.go
docker-dev-build:
	cp $(DEVDOCKER) Dockerfile
	$(BUILDIMG) ./ -t miraikeitai2020/bff:0.2.0
	rm Dockerfile
docker-dev-run:
	$(RUNIMG) -d -p 9020:9020 miraikeitai2020/bff:0.2.0
