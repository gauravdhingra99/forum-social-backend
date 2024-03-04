ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
UNIT_TEST_PACKAGES=$(shell  go list ./... | grep -v "vendor")
APP_EXECUTABLE="out/social_forum"

DB_NAME="social_forum"
DB_PORT=5432
TEST_DB_NAME="social_forum"
TEST_DB_PORT=5432

COVERAGE_MIN=70

GO111MODULE=on
GOPROXY=https://proxy.golang.org,direct
GOSUMDB=sum.golang.org

export GO111MODULE
export GOPROXY
export GOSUMDB

.PHONY: all
all: fmt vet lint build test

build:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE) ./cmd/*.go

install:
	go install ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

fmtcheck:
	@gofmt -l -s $(SOURCE_DIRS) | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

lint:
	@if [[ `golint $(ALL_PACKAGES) | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } | wc -l | tr -d ' '` -ne 0 ]]; then \
          golint $(ALL_PACKAGES) | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; }; \
          exit 2; \
    fi;

test: testdb.reset
	@ENVIRONMENT=test go test  -race $(UNIT_TEST_PACKAGES)

test-ci: build copy-ci-config testdb.migrate
	ENVIRONMENT=test go test -race -p=1 -cover -coverprofile=coverage-temp.out $(UNIT_TEST_PACKAGES)
	@cat ./coverage-temp.out | grep -v "mock.go" > ./coverage.out
	@go tool cover -func=coverage.out
	@go tool cover -func=coverage.out | gawk '/total:.*statements/ {if (strtonum($$3) < $(COVERAGE_MIN)) {print "ERR: coverage is lower than $(COVERAGE_MIN)"; exit 0}}'

test-cover-html:
	@echo "mode: count" > out/coverage-all.out
	$(foreach pkg, $(ALL_PACKAGES),\
	ENVIRONMENT=test go test -coverprofile=out/coverage.out -covermode=count $(pkg);\
	tail -n +2 out/coverage.out >> out/coverage-all.out;)
	go tool cover -html=out/coverage-all.out -o out/coverage.html

.PHONY: copy-config
copy-config:
	cp ./configs/application.sample.yml ./configs/application.yml
	cp ./configs/test.sample.yml ./configs/test.yml

.PHONY: copy-ci-config
copy-ci-config:
	cp ./configs/ci.sample.yml ./configs/test.yml


db.create:
	createdb -p $(DB_PORT) -Opostgres -Eutf8 $(DB_NAME)

db.migrate: build
	$(APP_EXECUTABLE) migrate

db.rollback: build
	$(APP_EXECUTABLE) rollback

db.drop:
	dropdb -p $(DB_PORT) --if-exists -Upostgres $(DB_NAME)

db.reset: db.drop db.create db.migrate

testdb.create:
	createdb  -p $(TEST_DB_PORT) -Opostgres -Eutf8 $(TEST_DB_NAME)

testdb.migrate: build
	ENVIRONMENT=test $(APP_EXECUTABLE) migrate

testdb.drop:
	dropdb -p $(TEST_DB_PORT) --if-exists -Upostgres $(TEST_DB_NAME)

testdb.reset: testdb.drop testdb.create testdb.migrate
