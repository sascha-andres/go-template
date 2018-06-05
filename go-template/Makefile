.PHONY: test

test:
		-rm -rf test
		go run main.go new --name test --template go-rest --arguments Namespace=github.com/sascha-andres/go-template/go-template --log-level debug
		-cd test/test && go run main.go serve
