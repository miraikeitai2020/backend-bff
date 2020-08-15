GOCMD=go
GORUN=$(GOCMD) run
GOGET=$(GOCMD) get
DEVCONFIG=config/development/gqlgen.yml
GQLGEN=github.com/99designs/gqlgen
TMPGEN= mkdir tmp

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
