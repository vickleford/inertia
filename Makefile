all: build

.EXPORT_ALL_VARIABLES:
GO11MODULE=on

build:
	@echo "\033[92mINFO: Building inertia\033[0m"
	@go build -mod=vendor

test:
	@echo "\033[92mINFO: Running tests\033[0m"
	@find . -name '*_test.go' -exec dirname {} \; | xargs -L1 bash -c 'cd "$$0" && go test .'
