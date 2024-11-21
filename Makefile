setup:
	go install github.com/mgechev/revive@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest

lintAll:
	go vet ./...
	go fmt ./...
	revive -config reviveConfig.toml -formatter friendly ./...

lint:
	go vet ${FILENAME}
	go fmt ${FILENAME}
	revive -config reviveConfig.toml -formatter friendly ${FILENAME}
