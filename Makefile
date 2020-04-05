default:
	go generate .
	go build .

test: default
	cd example && make test