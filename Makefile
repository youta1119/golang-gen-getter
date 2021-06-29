

.PHONY: e2e_test

e2e_test:
	rm -rf e2e_test/*_gen.go
	go generate e2e_test/*.go
	go test e2e_test/*.go