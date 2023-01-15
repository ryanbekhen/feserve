DEV_TOOLS = github.com/cosmtrek/air@latest \
	github.com/goreleaser/goreleaser@latest

install:
	@echo "installing dev tools..."
	@for i in $(DEV_TOOLS); do \
		go install $$i;	\
	done
	
	@echo "installing dependencies..."
	@go mod download

start-dev:
	@air

start:
	@go run .

build: clean
	@echo "compiling..."
	@go build -o bin/feserve .

clean:
	@rm -rf bin/ tmp/