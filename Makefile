lint:
	golint **/*.go -min_confidence 0.1 -set_exit_status

test:
	go test